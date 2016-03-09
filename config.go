package main

import (
  "github.com/wpxiong/beargo/filter"
)

func InitConfig() map[string]filter.FilterFunc {
  funcMap := make(map[string]filter.FilterFunc)
  
  //start
  funcMap["ParameterParseFilter"] = filter.ParameterParseFilter
  funcMap["ParameterBinderFilter"] =  filter.ParameterBinderFilter
  funcMap["RenderBindFilter"] =  filter.RenderBindFilter
  funcMap["RenderOutPutFilter"] =  filter.RenderOutPutFilter
  //end
  
  return funcMap
}