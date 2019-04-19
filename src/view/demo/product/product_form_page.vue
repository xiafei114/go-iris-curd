<template>
  <div>
    <Form ref="entityForm" :model="entity" :rules="ruleValidate" label-position="right" :label-width="110">
      <Row :gutter="16">
        <h4>基本信息</h4>
        <FormItem prop="productCode" label="产品编号：">
          <Input type="text" v-model="entity.productCode" placeholder="Enter text" clearable
                 style="width: auto"/>
        </FormItem>
        <FormItem prop="productCateName" label="产品类别：">
          <Input type="text" v-model="entity.productCateName" placeholder="Enter text" clearable
               style="width: auto" readonly/>
          <input type="hidden" v-model="entity.productCateId"></input>
          <Button class="tool" type="primary" icon="md-person-add" style="margin-left: 5px" @click="showProductCateModal">选择类别</Button>
        </FormItem>

        <FormItem prop="productName" label="产品名称：">
          <Input type="text" v-model="entity.productName" placeholder="Enter text" clearable
                 style="width: auto"/>
        </FormItem>

        <FormItem prop="productType" label="产品类型：">
          <Select v-model="entity.productType" style="width:200px">
            <Option v-for="item in productTypeList" :value="item.value" :key="item.value">{{ item.label }}</Option>
          </Select>
        </FormItem>

        <FormItem prop="productColor" label="产品颜色：">
          <RadioGroup v-model="entity.productColor">
            <Radio label="红"></Radio>
            <Radio label="黄"></Radio>
            <Radio label="蓝"></Radio>
          </RadioGroup>
        </FormItem>

        <FormItem prop="price" label="单价：">
          <Input type="text" v-model="entity.price" number placeholder="Enter text" clearable
                 style="width: auto"/>
        </FormItem>
        <FormItem prop="number" label="数量：">
          <Input type="text" v-model="entity.number" number placeholder="Enter text" clearable
                 style="width: auto"/>
        </FormItem>
        <FormItem prop="startDate" label="开始时间：">
          <DatePicker type="date" style="width: 200px" v-model="entity.startDate" @on-change="onChangeDate" format="yyyy-MM-dd" placeholder="选择日期以及时间"></DatePicker>

        </FormItem>
        <FormItem prop="isValid" label="数量：">
          <Checkbox v-model="entity.isValid" style="width: auto">是否可用</Checkbox>
        </FormItem>




      </Row>

      <FormItem>
        <Button type="success" icon="md-checkmark" @click="handleSubmit('entityForm')">保存</Button>
        <Button type="error" icon="md-return-left" style="margin-left: 8px" @click="handleReset('entityForm')">重置
        </Button>
        <Button type="warning" icon="md-close" style="margin-left: 8px" @click="handleCloseTag">关闭</Button>
      </FormItem>

    </Form>

    <PopProductCateListModals ref="popProductCateList" @on-select="handleSelect"></PopProductCateListModals>
  </div>
</template>

<script>
  import {mapMutations} from 'vuex'
  import {Get,Post,Put} from '@/api/data'

  import PopProductCateListModals from '../../common/pop_product_cate_list_modals'

  export default {
    components: {
      PopProductCateListModals,
    },
    data() {
      return {
        entityBaseUrl:'/product',
        entityFormName:'product_form_page',
        isEdit:false,
        entity: {
          productCode:'',
          productName:'',
          productType:'',
          productColor:'',
          price:0,
          number:0,
          productCateName:'',
          productCateId:'',
          isValid:true,
          startDate:'',
        },
        ruleValidate: {
          productCode: [
            {required: true, trigger: 'blur',max: 20 ,min:1}
          ],
          productName: [
            {required: true, message: '名称不能为空', trigger: 'blur',max: 50 ,min:1}
          ],
          productCateName: [
            {required: true, message: '类别不能为空', trigger: 'blur',max: 50 ,min:1}
          ],
          productType: [
            {required: true, message: '类型不能为空', trigger: 'blur'}
          ],
          productColor: [
            {required: true, message: '颜色需要选择', trigger: 'blur'}
          ],
          startDate: [
            {
              type:"date",
              required: true, message: '开始时间需要选择', trigger: 'blur'}
          ],


        },
        productTypeList: [
          {
            value: 'A',
            label: 'A'
          },
          {
            value: 'B',
            label: 'B'
          },
          {
            value: 'C',
            label: 'C'
          }
        ]
      }
    },
    methods: {
      ...mapMutations([
        'closeTag'
      ]),
      async getData() {
        let id = this.$route.params.id;
        if (id === undefined) {
          return;
        }

        let url = this.entityBaseUrl + '/' + id

        console.log(url);
        let resp = await Get(url);

        console.log(resp);
        this.entity = resp.data.data;
        this.isEdit = true;

      },
      handleSubmit(name) {
        console.log(this.entity)
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Loading.start();
            if (this.isEdit) {
              let url = this.entityBaseUrl;
              Put(url, this.entity).then(resp => {
                this.$Loading.finish();
                this.$Message.success('Success!')
                this.handleCloseTag()
              })
            } else {
              let url = this.entityBaseUrl
              Post(url, this.entity).then(resp => {
                this.$Loading.finish();
                this.$Message.success('Success!')
                this.handleCloseTag()
              })
            }
          } else {
            this.$Message.error('Fail!')
          }
        })
      },
      handleReset(name) {
        if (this.isEdit) {
          this.getData();
        } else {
          this.$refs[name].resetFields()
        }
      },
      handleCloseTag() {
        if (this.isEdit) {
          this.closeTag({
            name: this.entityFormName,
            params: {id: this.$route.params.id}
          })
        } else {
          this.closeTag({
            name: this.entityFormName
          })
        }
      },
      // 显示新建部门modal
      showProductCateModal () {
        this.$refs.popProductCateList.showPopModal(true)
      },
      handleSelect (cate) {
        console.log(cate.chnName,cate.id);
        this.entity.productCateName = cate.chnName;
        this.entity.productCateId = cate.id;
        this.$refs.popProductCateList.showPopModal(false)
      },
      onChangeDate(date){
      },

    },
    watch: {
      entity: {
        deep: true
      }
    },
    computed: {},
    created() {
      this.getData();
    },
    mounted: function () {

    }
  }
</script>
<style scoped>
  h4 {
    font-weight: 500;
    color: #6379bb;
    font-size: 15px;
    border-bottom: 1px solid #ddd;
    margin: 8px 10px 25px 10px;
    padding-bottom: 5px;
  }
</style>
