import axios from 'axios'
import router from '@/router'
import { getToken } from '@/libs/util'
// import { Spin } from 'iview'
import { Message } from 'iview'
const addErrorLog = errorInfo => {
  const { status } = errorInfo
  // 添加提示
  let tip = {
    content: errorInfo.data.msg,
    duration: 6,
    closable: true
  }
  switch (status) {
    case 200: Message.success(tip); break
    case 401: Message.warning(tip); router.push({ name: 'login' }); break // 回话过期
    case 403: Message.warning(tip); break
    case 404: Message.warning(tip); break
    case 500: Message.error(tip.content = 'Oh~~鬼知道服务器经历了什么~'); break
    case 501: Message.error(tip.content = '网关连接超时~'); break
    default: Message.error(tip.content = ('错误, 状态码[' + status + ']')); break
  }
}

class HttpRequest {
  constructor (baseUrl = baseURL) {
    this.baseUrl = baseUrl
    this.queue = {}
  }

  getInsideConfig () {
    const config = {
      baseURL: this.baseUrl,
      headers: {
        //
      }
    }
    return config
  }

  destroy (url) {
    delete this.queue[url]
    if (!Object.keys(this.queue).length) {
      // Spin.hide()
    }
  }

  interceptors (instance, url) {
    // 请求拦截
    instance.interceptors.request.use(config => {
      // 添加全局的loading...
      if (!Object.keys(this.queue).length) {
        // Spin.show() // 不建议开启，因为界面不友好
      }
      this.queue[url] = true
      config.headers['Authorization'] = 'bearer ' + getToken()
      return config
    }, error => {
      return Promise.reject(error)
    })
    // 响应拦截
    instance.interceptors.response.use(res => {
      this.destroy(url)
      const { data, status } = res
      return { data, status }
    }, error => {
      this.destroy(url)
      let errorInfo = error.response
      if (!errorInfo) {
        const { request: { statusText, status }, config } = JSON.parse(JSON.stringify(error))
        errorInfo = {
          statusText,
          status,
          request: { responseURL: config.url }
        }
      }
      addErrorLog(errorInfo)
      return Promise.reject(error)
    })
  }

  request (options) {
    const instance = axios.create()
    options = Object.assign(this.getInsideConfig(), options)
    this.interceptors(instance, options.url)
    return instance(options)
  }
}

export default HttpRequest
