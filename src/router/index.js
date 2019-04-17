import Vue from 'vue'
import Router from 'vue-router'
import routes from './routers'
import store from '@/store'
import iView from 'iview'
import { setToken, getToken, canTurnTo, setTitle } from '@/libs/util'
import { initMenu } from '@/api/routers'
// import config from '@/config'
// const { homeName } = config

Vue.use(Router)
export const router = new Router({
  routes
  // mode: 'history'
})
const LOGIN_PAGE_NAME = 'login'

const turnTo = (to, access, next) => {
  console.log('turnTo-->', to.name)
  if (canTurnTo(to.name, access, routes)) { // 有权限，可访问
    console.log('有权限，可访问 ***')
    // 加载动态菜单
    console.log('加载动态菜单-->2')
    let status = initMenu(router, store)
    console.log('************///////////// init menu, return status=', status)
    // if (status === 200) {
    //   next()
    // }else {
    //   next({ replace: true, name: 'login' })
    // }
    next()
  } else { // 无权限，重定向到401页面
    console.log('无权限，重定向到登录页面 ***')
    // next({ replace: true, name: 'error_401' })
    next({ replace: true, name: 'login' })
  }
}

router.beforeEach((to, from, next) => {
  iView.LoadingBar.start()
  const token = getToken()
  let has = window.sessionStorage.getItem('hasGetInfo')
  console.log('to.name=', to.name, ' from.name=', from.name)
  console.log('1------------', token, has)

  // 直接要跳登录页
  if (to.name === LOGIN_PAGE_NAME) {
    console.log('直接要跳登录页')
    setToken('')
    next()
    return
  }

  console.log('2------------', token, has)
  // 没有用户信息
  if (!has) {
    console.log('只要没有用户信息, 跳转到登录页')
    setToken('')
    next({ name: LOGIN_PAGE_NAME })
    return
  }

  console.log('3------------', token, has)
  // 未登陆, 且要跳转登录页
  if (!token && to.name === LOGIN_PAGE_NAME) {
    console.log('未登陆, 且要跳转登录页')
    setToken('')
    next() // 跳转
    return
  }

  console.log('4------------', token, has)
  // 未登录, 且要跳转非登录页
  if (!token && to.name !== LOGIN_PAGE_NAME) {
    console.log('未登录, 且要跳转非登录页')
    setToken('')
    next({
      name: LOGIN_PAGE_NAME // 跳转到登录页
    })
    return
  }

  console.log('to name=', to.name, ', from name=', from.name, ', to path=', to.path, ', from path=', from.path)
  if (to.name === null) {
    next({ path: '/' })
    return
  }
  turnTo(to, store.state.user.access, next)
  console.log('end------------', token, has)

  //
  // // 只要没有用户信息, 跳转到登录页
  // if (!has) {
  //   console.log('只要没有用户信息, 跳转到登录页')
  //   setToken('')
  //   next({name: LOGIN_PAGE_NAME})
  //   return
  // }
  //
  // // 已登录, 且要跳转登录页
  // if (token && has && to.name === LOGIN_PAGE_NAME) {
  //   console.log('已登录、有用户信息, 且要跳转登录页 ***')
  //   initMenu(router, store)
  //   next({name: homeName})
  //   return
  // }
  //
  // // 已登录, 且要跳转非登录页
  // if (token && to.name !== LOGIN_PAGE_NAME) {
  //   console.log('已登录、有用户信息, 且要跳转登录页 ***')
  //   // initMenu(router, store)
  //   turnTo(to, store.state.users.access, next)
  // }

  // if (!token && to.name !== LOGIN_PAGE_NAME) {
  //   console.log('未登录且要跳转的页面不是登录页 ***')
  //   // 未登录且要跳转的页面不是登录页
  //   next({
  //     name: LOGIN_PAGE_NAME // 跳转到登录页
  //   })
  // } else if (!token && to.name === LOGIN_PAGE_NAME) {
  //   console.log('未登陆且要跳转的页面是登录页 ***')
  //   // 未登陆且要跳转的页面是登录页
  //   next() // 跳转
  // } else if (token && to.name === LOGIN_PAGE_NAME) {
  //   console.log('已登录且要跳转的页面是登录页 ***')
  //   console.log('------initMenu')
  //   // 加载动态菜单
  //   initMenu(router, store)
  //   // 已登录且要跳转的页面是登录页
  //   next({
  //     name: homeName // 跳转到homeName页
  //   })
  // } else {
  //   console.log('else ***, ', store.state.users.hasGetInfo)
  //   turnTo(to, store.state.users.access, next)
  //   if (store.state.users.hasGetInfo) {
  //     turnTo(to, store.state.users.access, next)
  //   } else {
  //     turnTo(to, store.state.users.access, next)
  //     setToken('')
  //     next({
  //       name: 'login'
  //     })
  //   }}

  // if (store.state.users.hasGetInfo) {
  //   turnTo(to, store.state.users.access, next)
  // } else {
  //   store.dispatch('getUserInfo').then(users => {
  //     // 拉取用户信息，通过用户权限和跳转的页面的name来判断是否有权限访问;access必须是一个数组，如：['super_admin'] ['super_admin', 'admin']
  //     turnTo(to, users.access, next)
  //   }).catch(() => {
  //     setToken('')
  //     next({
  //       name: 'login'
  //     })
  //   })
  // }
})

router.afterEach(to => {
  setTitle(to, router.app)
  iView.LoadingBar.finish()
  window.scrollTo(0, 0)
})

export default router
