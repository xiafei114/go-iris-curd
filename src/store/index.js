import Vue from 'vue'
import Vuex from 'vuex'

import user from './module/user'
import app from './module/app'
import {Ws} from './iris-ws.js';

Vue.use(Vuex)

const now = new Date();
const store = new Vuex.Store({
  state: {
    // 会话列表
    groups: [{
      group: {
        id: 'message',
        name: '我们热爱工作',
      },
      messages: [{
        user: {
          id: '1',
          name: 'kwokcc',
          avatar: 'dist/images/wa.png',
        },
        content: '欢迎来到王者荣耀！',
        date: now
      }]
    }

    ],
    //是否登录
    isLogin: false,
    //websocket连接
    socket: '',
    // 过滤出只包含这个key的会话
    filterKey: ''
  },
  mutations: {
    //连接websocket
    connection(state,callback) {
      console.log("do socket")
      state.socket = new Ws('ws://localhost:8088/chat');

      //连接建立
      state.socket.OnConnect(function () {
        state.socket.Emit("join", {
          "myRooms": 'message'
        });
        console.log('socket 加入成功');
      });

      //连接丢失
      state.socket.OnDisconnect(function () {
        // alert('连接丢失');
        console.log('socket 连接丢失');
      });

      //接收聊天消息
      state.socket.On("chat", function (msg) {
        console.log(msg);
        let message = JSON.parse(msg);
        console.log("接收到消息:",message);

        let group = state.groups.find(item => item.group.id === 'message');
        group.messages.push({
          user: {
            id: '',
            name: '',
            avatar: '',
          },
          content: message.content,
          date: new Date(),
          self: false
        });
        callback();
      });
    },

    // 发送消息
    sendMessage({
                  socket
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

store.watch(
  (state) => state.groups,
  (val) => {
    console.log('CHANGE: ', val);
  }, {
    deep: true
  }
);

export default store;
