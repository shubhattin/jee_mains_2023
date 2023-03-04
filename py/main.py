from fastapi import FastAPI, Request
from fastapi.responses import Response
from fastapi.middleware.cors import CORSMiddleware
from brotli_asgi import BrotliMiddleware
from datetime import timedelta
from kry.datt import DEV_ENV
import mains
import os

APP_NAME = "JEE Mains 2023 Score Calculator"
if DEV_ENV:
    app = FastAPI(debug=True, title=APP_NAME)
else:
    app = FastAPI(openapi_url=None, redoc_url=None, title=APP_NAME)

if DEV_ENV:  # Currently not working in the production(deta)
    app.add_middleware(BrotliMiddleware)
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
                "X-Robots-Tag": "noindex",
                "X-Frame-Options": "deny",
                "Cache-Control": "No-Store" if DEV_ENV else "public, max-age=0",
                # Using the E-tag caching instead
            }
        )
    for x in head:
        res.headers[x] = head[x]
    return res


app.include_router(mains.router)

# from kry.plugins import sthaitik_sanchit
# if PROD_ENV:
#     app.mount("/", sthaitik_sanchit(directory="public"), name="static")

if DEV_ENV:
    import uvicorn

    if __name__ == "__main__":
        uvicorn.run("main:app", host="0.0.0.0", port=3428, reload=True)
