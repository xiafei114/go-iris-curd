import HttpRequest from '@/libs/axios'
import config from '@/config'

const baseUrl = process.env.NODE_ENV === 'dev' ? config.baseUrl.dev : config.baseUrl.pro
const axios = new HttpRequest(baseUrl)
export default axios
