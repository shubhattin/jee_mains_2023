from fastapi import FastAPI, Request
from fastapi.responses import Response
from datetime import timedelta
from kry.datt import DEV_ENV, PROD_ENV
from kry.plugins import sthaitik_sanchit
import mains
from fastapi.middleware.cors import CORSMiddleware
import os

APP_NAME = "JEE Mains 2023 Score Calculator"
if DEV_ENV:
    app = FastAPI(debug=True, title=APP_NAME)
else:
    app = FastAPI(openapi_url=None, redoc_url=None, title=APP_NAME)

app.add_middleware(
    CORSMiddleware,
    allow_origins=[
        os.getenv("STATIC_SITE_URL"),
    ],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
CACHE_DURATION = int(timedelta(hours=8).total_seconds())


@app.middleware("http")
async def middleware(req: Request, call_next):
    res: Response = await call_next(req)
    head = {"X-sraShTA": "bhagavatprasAdAt"}
    if req.method == "GET":
        head.update(
            {
                "X-Robots-Tag": "index",
                "X-Frame-Options": "deny",
                "Cache-Control": "No-Store"
                if DEV_ENV
                else f"public, max-age={CACHE_DURATION}",
            }
        )
    for x in head:
        res.headers[x] = head[x]
    for x in ["X-Powered-By"]:
        if x in res.headers:
            del res.headers[x]
    return res


app.include_router(mains.router)

if PROD_ENV:
    app.mount("/", sthaitik_sanchit(directory="public"), name="static")

if DEV_ENV:
    import uvicorn

    if __name__ == "__main__":
        uvicorn.run("main:app", host="0.0.0.0", port=3428, reload=True)