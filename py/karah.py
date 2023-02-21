import shubhlipi as sh

for x in sh.argv:
    if x == "requirements.txt":
        sh.cmd(
            "pipenv requirements --exclude-markers > requirements.txt", display=False
        )
        fl = sh.read("requirements.txt")
        fl = fl.replace("-i https://pypi.org/simple\n", "")
        sh.write("requirements.txt", fl)
    elif x == "build":
        sh.cmd("pnpm build", direct=False)
        sh.copy_folder("../build", "public")
        print("\nCopied SvelteKit 'build' folder to FastAPI 'public' folder")
