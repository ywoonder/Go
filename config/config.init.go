package config

import(
  "log"
  "gopkg.in/gcfg.v1"
  "os"
)

type Config struct{

 Server ServerConfig
 ShopDB DBConfig
}

func ReadModuleConfig(cfg interface{}, path string, module string) bool {
  environ := os.Getenv("TKPENV")
  if environ == "" {
    environ = "development"
  }
  fname := path + "/" + module + "." + environ + ".ini"
  err := gcfg.ReadFileInto(cfg, fname)
  if err == nil {
    println("read config from ", fname)
    return true
  }
  log.Println(err)
  return false
}

var cfg *Config

func InitConfig (path ...string) *Config{
  config := Config{}
  for _,p := range path{
    ok := ReadModuleConfig(&config, p, "")
    if ok{
      cfg = &config
    }
    fmt.Println("failed to read config")
  }
  return nil
}
