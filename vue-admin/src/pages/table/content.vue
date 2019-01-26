<template>
    <div class="panel">
        <panel-title :title="$route.meta.title">
            <!--<el-button @click.stop="on_refresh" size="small">-->
            <!--<i class="fa fa-refresh"></i>-->
            <!--</el-button>-->
            <!--<router-link :to="{name: 'tableAdd'}" tag="span">-->
            <!--<el-button type="primary" icon="plus" size="small">添加数据</el-button>-->
            <!--</router-link>-->
        </panel-title>
        <div class="panel-body">
            
            
            <el-row>
                <el-button @click.stop="get_table_data" size="small" type="primary">
                    <i class="el-icon-refresh"></i>
                    <span type="text">刷新</span>
                </el-button>
                
                <el-button @click.stop="showPost" size="small" type="primary">
                    <i class="el-icon-plus "></i><span type="text">创建</span>
                </el-button>
                <br>
                <br>
            </el-row>
            
            
            <el-table
                :data="table_data"
                v-loading="load_data"
                element-loading-text="拼命加载中"
                border
                @selection-change="on_batch_select"
                style="width: 100%;">
                
                
                
                <!--<el-table-column prop="id" label="userid" width=""></el-table-column>-->
                <el-table-column prop="username" label="username" width=""></el-table-column>
                <!--<el-table-column prop="content.id" label="contentid" width=""></el-table-column>-->
                <el-table-column prop="content.title" label="title" width=""></el-table-column>
                <el-table-column prop="content.content" label="content" width=""></el-table-column>
                <!--<el-table-column prop="content" label="content" width=""></el-table-column>-->
                <!--<el-table-column prop="sex" label="性别" width="">-->
                <!--<template slot-scope="scope">-->
                <!--<span v-text="scope.row.sex == 1 ? '男' : '女'"></span>-->
                <!--</template>-->
                <!--</el-table-column>-->
                
                
                <el-table-column prop="content.createTime" label="createTime" width="200"></el-table-column>
                <el-table-column prop="content.updateTime" label="updateTime" width="200"></el-table-column>
                
                
                <el-table-column label="操作" width="180">
                    <template slot-scope="scope">
                        <!--<router-link :to="{name: 'tableUpdate', params: {id: scope.row.id}}" tag="span">-->
                        <!--<el-button type="info" size="small" icon="edit">修改</el-button>-->
                        <!--</router-link>-->
                        
                        
                        <!--
                                                           @click="showViewDialog(scope.row)"

                        -->
                        <el-button type="info" size="small"
                        
                                   @click="showPut(scope.row)"
                                   icon="edit">修改
                        </el-button>
                        
                        <el-button type="danger" size="small" icon="delete" @click="delete_data(scope.row)">删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            
            <el-pagination
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-size="10"
                layout="total, prev, pager, next"
                :total="total">
            </el-pagination>
            
            
            <!--<bottom-tool-bar>-->
            <!--<el-button-->
            <!--type="danger"-->
            <!--icon="delete"-->
            <!--size="small"-->
            <!--:disabled="batch_select.length === 0"-->
            <!--@click="on_batch_del"-->
            <!--slot="handler">-->
            <!--<span>批量删除</span>-->
            <!--</el-button>-->
            <!--<div slot="page">-->
            <!--<el-pagination-->
            <!--@current-change="handleCurrentChange"-->
            <!--:current-page="currentPage"-->
            <!--:page-size="10"-->
            <!--layout="total, prev, pager, next"-->
            <!--:total="total">-->
            <!--</el-pagination>-->
            <!--</div>-->
            <!--</bottom-tool-bar>-->
        
        
        </div>
        
        
        <el-dialog :title="dialogTypeName" :visible.sync="dialogFormVisible">
            <el-form :model="form" :rules="rules" ref="form" label-width="100px">
                
                <el-form-item label="主题" prop="">
                    <el-input v-model.trim="form.content.title"></el-input>
                </el-form-item>
                
                <el-form-item label="主要内容" prop="content">
                    <el-input type="textarea"
                              :autosize="{ minRows: 5,}"
                              v-model="form.content.content"></el-input>
                </el-form-item>
            </el-form>
            
            <div slot="footer" class="dialog-footer" v-if="dialogType === 1 || dialogType === 2">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <!--<el-button type="primary" @click="postOrPut()">确 定</el-button>-->
                <el-button type="primary" @click="submitForm('form')">确 定</el-button>
            
            </div>
        
        </el-dialog>
        
        
        <!--<p>total :{{total}}</p>-->
        <!--<p>sessionbean :{{sessionbean}}</p>-->
        <!--<p>form :{{form}}</p>-->
    
    </div>
</template>
<script type="text/javascript">
    import {bottomToolBar, panelTitle} from 'components'
    import global_ from '@/tool/Global'


    const url = global_.urlPrefix + "/content"


    export default {
        data() {
            return {
                table_data: null,
                //当前页码
                currentPage: 1,
                //数据总条目
                total: 0,
                //每页显示多少条数据
                length: 15,
                //请求时的loading效果
                load_data: true,
                //批量选择数组
                batch_select: [],


                dialogFormVisible: false,
                dialogTypeName: '修改',  //创建
                dialogType: 1,


                form: {
                    content: {}
                },

                rules: {
                    // title: [
                    //     {required: true, message: '不能为空', trigger: 'blur'},
                    //     {type: 'string', message: '长度需要在1~50之间', min: 1, max: 50}
                    // ],
                    // content: [
                    //     {required: true, message: '不能为空', trigger: 'blur'},
                    //     {type: 'string', message: '长度需要在1~400之间', min: 1, max: 400}
                    // ],

                },

                sessionbean: null,

            }
        },
        components: {
            panelTitle,
            bottomToolBar
        },
        created() {
            this.get_table_data()
        },
        methods: {
            //刷新
            on_refresh() {
                this.get_table_data()
            },
            //获取数据
            get_table_data() {
                // this.load_data = true
                // this.$fetch.api_table.list({
                //     page: this.currentPage,
                //     length: this.length
                // })
                //     .then(({data: {result, page, total}}) => {
                //         this.table_data = result
                //         this.currentPage = page
                //         this.total = total
                //         this.load_data = false
                //     })
                //     .catch(() => {
                //         this.load_data = false
                //     })

                this.get_method(this.currentPage)
                // this.session()

            },

            get_method(pagenum) {
                this.load_data = true
                // let this = this
                this.$axios.get(url + '?pageNum=' + pagenum + '&pageSize=' + global_.pageSizeString)
                    .then((response) => {
                        this.load_data = false
                        if (response.data.success === true) {
                            this.table_data = response.data.result
                            // this.pageMap = response.data.page
                            this.total = response.data.page.count
                        }

                        global_.handleResponseGetThen(response)
                    })
                    .catch((error) => {
                        this.load_data = false
                        this.$message.error(error.response.data.message);
                    })
            },

            session() {

                this.$axios.get(global_.urlPrefix + "/session")
                    .then((response) => {
                        this.sessionbean = response.data.result

                        // this.set_user_session(session)
                        // if (this.user_session.isLogin === true) {
                        //     this.$router.push({name: 'partyActivity'})
                        // }
                        // if (this.user_session.isLogin !== true) {
                        //
                        // }
                    })
                    .catch((error) => {
                    })
            },


            showPost(data) {
                this.dialogType = 1
                this.dialogTypeName = '创建'
                this.form = {
                    title: "",
                    content: {},
                }

                this.dialogFormVisible = true

            },

            showPut(data) {
                this.dialogType = 2
                this.dialogTypeName = '修改'

                this.$axios.get(url + '/' + data.content.id)
                    .then((response) => {
                        if (response.data.success === true) {
                            this.form = JSON.parse(JSON.stringify(response.data.result))
                            this.dialogFormVisible = true
                        }
                    })
                    .catch((error) => {
                    })
            },

            postOrPut() {
                this.dialogFormVisible = false;

//                post
                if (this.dialogType === 1) {
                    // this.form.userIds1 = this.userIds1
                    // this.form.userIds2 = this.userIds2
                    var str = JSON.stringify(this.form.content)

                    this.$axios.post(url, str)
                        .then((response) => {
                            global_.handleResponseThen(response)
                            // self.$message.success(response.data.message);
                            this.get_table_data()
                        })
                        .catch((error) => {
                            global_.handleResponseCatch()
                        })
                }

                // put
                if (this.dialogType === 2) {
                    var form = JSON.parse(JSON.stringify(this.form.content))
                    delete form.createTime
                    delete form.updateTime

                    this.$axios.put(url, JSON.stringify(form))
                        .then((response) => {
                            global_.handleResponseThen(response)
                            this.get_table_data()
                            // self.$message.success(response.data.message);
                        })
                        .catch((error) => {
                            global_.handleResponseCatch()
                        })
                }
            },

            submitForm(formName) {
                let self = this
                this.$refs[formName].validate((valid) => {
                    if (valid) {

                        self.postOrPut()

                    } else {
                        return false;
                    }
                });
            },


            //单个删除
            delete_data(item) {
                this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                })
                    .then(() => {

                        this.$axios.delete(url + "/" + item.content.id)
                            .then((response) => {
                                global_.handleResponseThen(response)
                                this.get_table_data()
                                // self.$message.success(response.data.message);

                            })
                            .catch((error) => {
                                global_.handleResponseCatch()
                                // self.get_table_data()
                                //  self.$message.error(error.response.data.message);
                            })

                        // this.load_data = true
                        // this.$fetch.api_table.del(item)
                        //     .then(({msg}) => {
                        //         this.get_table_data()
                        //         this.$message.success(msg)
                        //     })
                        //     .catch(() => {
                        //     })


                    })
                    .catch(() => {
                    })
            },
            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val
                this.get_table_data()
            },
            //批量选择
            on_batch_select(val) {
                this.batch_select = val
            },
            //批量删除
            on_batch_del() {
                this.$confirm('此操作将批量删除选择数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                })
                    .then(() => {
                        this.load_data = true
                        this.$fetch.api_table.batch_del(this.batch_select)
                            .then(({msg}) => {
                                this.get_table_data()
                                this.$message.success(msg)
                            })
                            .catch(() => {
                            })
                    })
                    .catch(() => {
                    })
            }
        }
    }
</script>
