<template>

    <Form ref="userValidate" :model="userValidate" :rules="ruleValidate" label-position="right" :label-width="110">
      <Row :gutter="16">
        <Col span="12" push="2">
          <FormItem prop="name" label="部门名称：">
            <Input type="text" v-model="userValidate.name" placeholder="Enter name" clearable style="width: auto" />
          </FormItem>
          <FormItem prop="username" label="稷下之学：">
            <Input type="text" v-model="userValidate.username" prefix="ios-contact" placeholder="Enter text" clearable style="width: auto" />
          </FormItem>
          <FormItem prop="phone" label="手机号码：">
            <Input type="text" v-model="userValidate.phone" prefix="ios-phone-portrait" placeholder="Enter text" clearable style="width: auto" />
          </FormItem>
        </Col>

        <Col span="12" push="1">
          <FormItem label="邮箱：">
            <Input type="text" v-model="userValidate.email" prefix="ios-mail" placeholder="Enter text" clearable style="width: auto" />
          </FormItem>
          <FormItem prop="password"  label="登录密码：">
            <Input type="password" v-model="userValidate.password" prefix="ios-lock" placeholder="Enter text" clearable style="width: auto" />
          </FormItem>
        </Col>
      </Row>

      <FormItem :style="{padding: '0 400px'}">
        <Row>
          <Col span="8"><Button type="success" icon="md-checkmark" @click="handleSubmit('userValidate')">保存</Button></Col>
          <Col span="8"><Button type="error" icon="md-return-left" @click="handleReset('userValidate')">重置</Button></Col>
          <Col span="8"><Button type="warning" icon="md-close" @click="handleCloseTag">关闭</Button></Col>
        </Row>
      </FormItem>

    </Form>

</template>

<script>
import { mapMutations } from 'vuex'
export default {
  data () {
    return {
      genderSelect: [
        { label: '男', 'value': '男' },
        { label: '女', 'value': '女' },
        { label: '未知', 'value': '未知' }
      ],
      userValidate: {
        name: '',
        phone: '',
        username: '',
        gender: '',
        age: '',
        depId: '',
        email: '',
        password: '',
        enable: true
      },
      ruleValidate: {
        name: [
          { required: true, message: '名字不能为空', trigger: 'blur' }
        ],
        username: [
          { required: true, message: '登录账户不能为空', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '登录密码不能为空', trigger: 'blur' }
        ]
        // gender: [
        //   {required: true, message: '性别不能为空', trigger: 'change'}
        // ],
      }
    }
  },
  methods: {
    ...mapMutations([
      'closeTag'
    ]),
    handleSubmit (name) {
      this.$refs[name].validate((valid) => {
        if (valid) {
          this.$Message.success('Success!')
        } else {
          this.$Message.error('Fail!')
        }
      })
    },
    handleReset (name) {
      this.$refs[name].resetFields()
      this.userValidate = { enable: true }
    },
    handleCloseTag () {
      this.closeTag({
        name: 'add_user_page'
      })
    }
  },
  watch: {
  },
  computed: {},
  mounted: function () {
  }
}
</script>
<style scoped>
</style>
