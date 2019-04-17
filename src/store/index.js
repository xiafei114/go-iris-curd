import Vue from 'vue'
import Vuex from 'vuex'

import user from './module/user'
import app from './module/app'
import {Ws} from './iris-ws.js';

Vue.use(Vuex)

const now = new Date();
const store = new Vuex.Store({
  state: {
    //
    //是否登录
    isLogin: false,
    //websocket连接
    socket: '',
    // 过滤出只包含这个key的会话
    filterKey: ''
  },
  mutations: {
    //连接websocket
    connection(state) {
      console.log("do socket")
      state.socket = new Ws('ws://localhost:8088/chat');

      //连接建立
      state.socket.OnConnect(function () {
        state.socket.Emit("join", {
          "myRooms": 'message'
        });
        console.log('加入成功');
      });

      //连接丢失
      state.socket.OnDisconnect(function () {
        alert('连接丢失');
        console.log('连接丢失');
      });

      //接收聊天消息
      state.socket.On("chat", function (msg) {
        console.log(msg);
        let message = JSON.parse(msg);
        console.log(message);
      });
    },

    // 发送消息
    sendMessage({
                  user,
                  socket,
                  action
                }, content) {

      let sender = {
        id: user.id,
        name: user.name,
        avatar: user.img,
      };

      let msg = {
        user: sender,
        content: content
      };

      let message = {
        "roomId": 'message',
        "msg": JSON.stringify(msg)
      };
      socket.Emit("chat", message);
      console.log('消息已发送');
    },
  },
  actions: {
    //
  },
  modules: {
    user,
    app
  }
})


export default store;
