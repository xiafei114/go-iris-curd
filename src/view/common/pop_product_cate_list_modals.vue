<template>

  <div>
    <!--新增部门modal-->
    <Modal v-model="showModal" :z-index=1 draggable>
      <p slot="header">
        <Icon type="information-circled"></Icon>
        <span>选择类型</span>
      </p>
      <div>
        <tables ref="tables"
                searchable
                search-place="top"
                v-model="tableData"
                :columns="columns"
                :total="total"
                :highlightRow=true
                @change-page="pullTableList"
                @change-size="pullTableList"
                @handle-Search="pullTableList"
                @on-select="handleSelect"
                @on-selection-change="selectChange"
        />
      </div>

      <div slot="footer">
        <Button type="warning">重置</Button>
        <Button type="success" @click="close">关闭</Button>
      </div>
    </Modal>
  </div>

</template>

<script>
  import {Get} from '@/api/data'
  import Tables from '_c/simple-tables'
  import Button from "iview/src/components/button/button";

  export default {
    name: 'PopProductCateListModals',
    props: {},
    components: {
      Button,
      Tables,
    },
    data() {
      return {
        showModal: false,
        refreshEntityLoading: false,
        total: 0,
        columns: [
          {title: '编号', key: 'numCode',width: 100, sortable: true},
          {title: '名称', key: 'chnName', sortable: true},
          {
            title: '操作',
            key: 'handle',
            align: 'center',
            width: 100,
            button: [
              (h, params, vm) => {
                return h("div", [
                  h(
                    "Button",
                    {
                      props: {
                        type: "primary",
                        size: "small"
                      },
                      style: {
                        marginRight: "5px"
                      },
                      on: {
                        'click': () => {
                          vm.$emit('on-select', params)
                        }
                      }
                    },
                    "选中"
                  )
                ]);
              }
            ]
          }
        ],
        entityBaseUrl:'/productCate',
        tableData: [],
        selection:[],
        page: {
          start: 1,
          size: 10,
          searchKey :'',
          searchValue:''
        }
      }
    },
    methods: {
      handleSelect(params) {
        // console.log(params.row.id,params.row.chnName)
        this.$emit('on-select', params.row)
      },
      // 创建部门和修改部门的modal的状态操作
      showPopModal(show) {
        this.showModal = show
      },
      close() {
        this.showModal = false
      },
      // 监听选择的表格
      selectChange(selection) {

        if(selection instanceof Array){
          selection.length === 1 ? this.modifyEntityDisable = false : this.modifyEntityDisable = true;
          selection.length > 0 ? this.deleteEntitysDisable = false : this.deleteEntitysDisable = true;
          this.selection = selection;
        }
      },

      pullTableList(page) {
        this.tableData = []
        this.refreshEntityLoading = true

        this.page = page;

        if (this.page.start === undefined) this.page.start = 1
        if (this.page.size === undefined) this.page.size = 10
        let url = this.entityBaseUrl + '?start=' + this.page.start + '&size=' + this.page.size ;

        if(! (this.page.searchValue === undefined)){
          url +=  '&searchKey=' + this.page.searchKey + '&searchValue=' + this.page.searchValue
        }

        console.log('url=', url)
        // return
        Get(url).then(resp => {
          this.total = resp.data.total
          // resp.data.rows.forEach(d => this.tableData.push(d))
          this.tableData = resp.data.rows
        })
        setTimeout(() => {
          this.refreshEntityLoading = false
        }, 1.5 * 1000)
      },
      pullData() {
        this.pullTableList({})
      }
    },
    watch: {},
    computed: {},
    mounted() {
      this.pullData()
    }
  }
</script>
<style scoped>
</style>
