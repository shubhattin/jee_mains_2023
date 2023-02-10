from time import time
import shubhlipi as sh, json, yaml, os, re
import markdown, markdownify, jinja2

ln = sh.lang_list
ln2 = sh.dict_rev(ln)
if not os.path.isfile("r.json"):
    sh.write("r.json", sh.dump_json(dict(yes={}, no={}), 2))
if sh.args(0) in ln2:
    src = sh.args(0)
else:
    src = "en"
if input(f"Do you want to translate with {src} as base? ") != "yes":
    exit()

DEFAULT_LOCALE = "en"
sh.write(
    "locales.json",
    sh.dump_json(
        dict(locales=sh.dict_rev(sh.lang_list), default_locale=DEFAULT_LOCALE), 2
    ),
)
sh.prettier_beautify("locales.json")
main_db = yaml.safe_load(sh.read(f"data/{ln2[src]}.yaml"))
anu = {}


def render_from_string(template: str, **data):
    return (
        jinja2.Environment(loader=jinja2.BaseLoader)
        .from_string(template)
        .render(**data)
    )


TEMP = sh.read("template/md_template.md.j2")


def get_all_lang_markdown():
    res = {}
    MATCHES = re.compile(r"# .+?\n\n.+?\n\n", re.DOTALL)
    HEAD = re.compile(r"(?<=# ).+?(?=\n)")
    LEKH = re.compile(r"(?<=\n\n).+?(?=\n\n)", re.DOTALL)
    for lng in list(sh.lang_list_names.keys()):
        code = sh.lang_list[lng]
        res[code] = {}
        vl = sh.read(f"data/markdown/{lng}.md") + "\n"
        for x in MATCHES.findall(vl):
            head = HEAD.findall(x)[0]
            lekh = LEKH.findall(x)[0]
            res[code][head] = lekh
    return res


mark_db = get_all_lang_markdown()

only = json.loads(sh.read("r.json"))
print(only)


def pre_process(func=0):
    func = [sh.anuvadak, sh.parivartak][func]
    global mark_db

    def process_func(vl, src, to):
        vl = markdown.markdown(vl)
        if vl.startswith("<p>") and vl.endswith("</p>"):
            vl = vl[3:-4]
        template_found = re.match(r"^\(.+?\)$", vl) and vl[1:-1] in mark_db[to]
        tmp = vl
        templ_nm = None
        if template_found:
            templ_nm = vl[1:-1]
            vl = markdown.markdown(mark_db[src if to != "ur" else "hi"][templ_nm])
        data = func(vl, src, to, html=True)
        if template_found:
            mark_db[to][templ_nm] = markdownify.markdownify(data)
        return markdownify.markdownify(data) if not template_found else tmp

    return process_func


for y in ln:
    if ln[y] == "en":
        continue
    tm = time()
    anu[y] = {}
    org = yaml.safe_load(sh.read(f"data/{y}.yaml"))
    if ln[y] != "ur":
        anu[y] = sh.process_json(
            main_db,
            org,
            src,
            ln[y],
            no=only["no"],
            yes=only["yes"],
            only_org=ln[y] in ("sa", "hi"),
            func=pre_process(0),
        )
    else:
        anu[y] = sh.process_json(
            anu["हिन्दी"],
            org,
            "hi",
            "ur",
            no=only["no"],
            yes=only["yes"],
            func=pre_process(1),
        )
    sh.write(
        f"data/{y}.yaml",
        yaml.safe_dump(anu[y], allow_unicode=True, sort_keys=False),
    )
    sh.write(f"data/markdown/{y}.md", render_from_string(TEMP, DATA=mark_db[ln[y]]))

    print(ln[y], f"{time()-tm}s")

sh.start_thread(lambda: sh.prettier_beautify("data"))

if True:  # Adding a model for TypeScript
    model = main_db
    sh.write("model.ts", sh.generate_typescript_data_model(model, "dattType"))
    sh.start_thread(lambda: sh.prettier_beautify("model.ts"))
