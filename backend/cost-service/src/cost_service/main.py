"""Main entry point for the cost service."""

import logging
import os

from fastapi import FastAPI

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI(
    title="Hosterizer Cost Service",
    description="Cost management and tracking service",
    version="0.1.0",
)


@app.get("/health")
async def health_check() -> dict[str, str]:
    """Health check endpoint."""
    return {"status": "healthy"}


if __name__ == "__main__":
    import uvicorn

    port = int(os.getenv("PORT", "8007"))
    logger.info(f"Cost Service starting on port {port}...")
    uvicorn.run(app, host="0.0.0.0", port=port)
