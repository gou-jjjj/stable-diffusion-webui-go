"""
    1. Place this file in the root directory of stable-diffusion-webui.

    2. Set up the environment for stable-diffusion-webui (refer to: https://github.com/AUTOMATIC1111/stable-diffusion-webui) to ensure all dependencies are properly installed.

    3. After that, run this file. It will start a stable-diffusion-webui server with only the API and no frontend.
       We can then use Golang to connect to these APIs and interact with the server.
"""

from webui import api_only

api_only()
