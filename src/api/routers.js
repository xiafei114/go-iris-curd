import axios from '@/libs/api.request'
import Main from '@/components/main'


export const getRouterReq = (access) => {
  return axios.request({
    url: 'get_router',
    params: {
      access
    },
    method: 'get'
  })
}

/**
 * 初始化菜单
 * @param router
 * @param store
 */
export const initMenu = (router, store) => {
  console.log('****************************************  1', store.state.app.routes.length)
  console.log('****************************************  2', store.state.app.routes.length)
  if (store.state.app.routes.length > 0) return
  // TODO addRouter()，刷新了页面后 vuex和router里面的信息都被重新初始化了，此时要重新设置动态路由，  所以判断vuex中新增的数据，如果为空则用户刷新了页面
  // var menus = sessionStorage.getItem('menus')
  // console.log(JSON.parse(menus))
  // if (menus) {
  //   console.log(111)
  //   initRouterConfig(router, store, JSON.parse(menus))
  //   return
  // }
  let status = 0
  axios.request({
    url: 'sysMenu',
    method: 'get'
  }).then(resp => {
    console.log('req -->> /config/menus -------------------------')
    if (resp && resp.status === 200) {
      status = 200
      // 保存菜单
      // sessionStorage.setItem('menus', JSON.stringify(resp.data))
      initRouterConfig(router, store, resp.data.data)
    }
  })
  return status
}

// 动态添加路由
var initRouterConfig = function (router, store, menus) {
  var fmtRoutes = formatRoutes(menus)
  fmtRoutes.forEach(fr => {
    router.options.routes.push(fr)
    console.log(fr.name, fr.meta.access)
    if (fr.children instanceof Array) {
      fr.children.forEach(c => {
        console.log(c.name, c.meta.access)
      })
    }
    console.log('@@@@@@@@@@@@@')
  })
  router.addRoutes(router.options.routes) // 调用addRoutes添加路由
  console.log("do initRouterConfig")
  store.commit('initMenu', fmtRoutes)

}
// https://router.vuejs.org/zh/api/#router-构建选项
export const formatRoutes = (routes) => {
  let fmRoutes = []
  routes.forEach(router => {
    let {
      path,
      name,
      meta,
      children
    } = router
    if (children && children instanceof Array) {
      children = formatChildenRoutes(children)
    }

    let fmRouter = {
      path: path,
      component: Main,
      name: name,
      // icon: icon,
      meta: meta,
      children: children
    }
    fmRoutes.push(fmRouter)
  })
  return fmRoutes
}
const formatChildenRoutes = (childs) => {
  let fmRoutes = []
  childs.forEach(child => {
    let {
      path,
      name,
      modular,
      component,
      meta,
      children
    } = child

    if (children && children instanceof Array) {
      children.forEach(ch => {
        if (ch && ch instanceof Array) {
          children = formatChildenRoutes(ch)
        }
      })
    }

    let fmRouter = {
      path: path,
      component: (resolve) => {
        require(['@/view' + modular + component + '.vue'], resolve)
      },
      name: name,
      meta: meta,
      children: children
    }
    fmRoutes.push(fmRouter)
  })
  return fmRoutes
}
