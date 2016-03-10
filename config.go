package main

import (
  "github.com/wpxiong/beargo/filter"
  "github.com/wpxiong/beargo/session"
  "github.com/wpxiong/beargo/webapp"
  "github.com/wpxiong/beargo/render/template"
)

func InitConfig() webapp.ConfigMap{
  config := webapp.ConfigMap{}
  config.Filterfuncmap = InitFiltrConfig()
  config.Sessionprovidermap = InitSessionProviderConfig()
  config.Templatefuncmap = InitTemplatefuncConfig()
  return config
}

func InitFiltrConfig() map[string]filter.FilterFunc {
  funcMap := make(map[string]filter.FilterFunc)
  return funcMap
}


func InitSessionProviderConfig() map[string]session.SessionProvider {
  sessionProviderMap := make(map[string]session.SessionProvider)  
  return sessionProviderMap
}

func InitTemplatefuncConfig() template.TemplateFuncMap {
  templateFuncMap := make(template.TemplateFuncMap)  
  return templateFuncMap
}