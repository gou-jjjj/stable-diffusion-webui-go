## stable-diffusion-webui-go-client

If you want to use sdwgc, you must first have a [stable-diffusion-webui](https://github.com/AUTOMATIC1111/stable-diffusion-webui) service  .

1. Place [run.py](run.py) in the root directory of stable-diffusion-webui.

1. Set up the environment for stable-diffusion-webui (refer to: https://github.com/AUTOMATIC1111/stable-diffusion-webui) to ensure all dependencies are properly installed.

1. After that, run run.py. It will start a stable-diffusion-webui server with only the API and no frontend.
   We can then use Golang to connect to these APIs and interact with the server.

```shell 
    go get github.com/G-JJJJ/stable-diffusion-webui-go
```

Sample:
 ```go
package main

import "fmt"

func main() {
   sd := NewSD("http://127.0.0.1:7861", 20*time.Minute, true)
   
   img, err := sd.Txt2Img(DefaultTxt2ImgReq())
   if err != nil {
      fmt.Println(err)
      return
   }

   fmt.Println(img.ToString())
}
```

Everyone is welcome to provide the code and improve it ...

