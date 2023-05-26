package fuzz_trainbot_imutil

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"
    "bytes"
    "image"

    "github.com/jo-m/trainbot/pkg/imutil"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                img, _, err := image.Decode(bytes.NewReader(data))
                if err != nil {
                    return 0
                }
                
                imutil.Copy(img)
                return 0

            case 1:
                temp, _ := fuzzConsumer.GetInt()
                seed := int64(temp)
                w, _ := fuzzConsumer.GetInt()
                h, _ := fuzzConsumer.GetInt()

                imutil.RandGray(seed, w, h)
                return 0

            case 2:
                temp, _ := fuzzConsumer.GetInt()
                seed := int64(temp)
                w, _ := fuzzConsumer.GetInt()
                h, _ := fuzzConsumer.GetInt()

                imutil.RandRGBA(seed, w, h)
                return 0

            case 3:
                w, _ := fuzzConsumer.GetInt()
                h, _ := fuzzConsumer.GetInt()

                imutil.NewYuv420(data, w, h)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}