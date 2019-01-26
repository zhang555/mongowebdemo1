<template>
    <div class="panel">
        <panel-title :title="$route.meta.title"></panel-title>
        
        <div class="panel-body">
            <el-row>
                <el-button @click.stop="get_table_data" size="small" type="primary">
                    <i class="el-icon-refresh"></i>
                    <span type="text">刷新</span>
                </el-button>
                
                <br>
                <br>
            </el-row>
            
            
            <el-table
                :data="table_data"
                v-loading="load_data"
                element-loading-text="拼命加载中"
                border
                style="width: 100%;">
                
                
                <el-table-column prop="role.roleName" label="roleName" width=""></el-table-column>
                <el-table-column prop="role.roleNameShow" label="roleNameShow" width=""></el-table-column>
                <el-table-column prop="count" label="count" width=""></el-table-column>
                
                <el-table-column prop="usernames" label="usernames" width="">
                    <template slot-scope="scope">
                        
                        <p v-for=" value in scope.row.usernames">
                            {{value}}
                        </p>
                    </template>
                </el-table-column>
                
                
                <el-table-column label="操作" width="180">
                    <template slot-scope="scope">
                        
                        <el-button type="info" size="small" @click="showPut(scope.row)" icon="edit">修改</el-button>
                    
                    </template>
                </el-table-column>
            </el-table>
            
            <!--<el-pagination-->
            <!--@current-change="handleCurrentChange"-->
            <!--:current-page="currentPage"-->
            <!--:page-size="10"-->
            <!--layout="total, prev, pager, next"-->
            <!--:total="total">-->
            <!--</el-pagination>-->
        
        </div>
        
        
        <el-dialog :title="dialogTypeName" :visible.sync="dialogFormVisible">
            <el-form :model="form" :rules="rules" ref="form" label-width="120px">
                
                <el-form-item label="roleName" prop="roleName">
                    <el-input v-model.trim="form.role.roleName" disabled></el-input>
                </el-form-item>
                
                <el-form-item label="roleNameShow" prop="roleNameShow">
                    <el-input v-model.trim="form.role.roleNameShow"></el-input>
                
                </el-form-item>
            </el-form>
            
            <div slot="footer" class="dialog-footer" v-if="dialogType === 1 || dialogType === 2">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <!--<el-button type="primary" @click="postOrPut()">确 定</el-button>-->
                <el-button type="primary" @click="submitForm('form')">确 定</el-button>
            
            </div>
        
        </el-dialog>
        
        
        <!--<p>total :{{total}}</p>-->
        <!--<p>table_data :{{table_data}}</p>-->
    
    </div>
</template>
<script type="text/javascript">
    import {bottomToolBar, panelTitle} from 'components'
    import global_ from '@/tool/Global'


    const url = global_.urlPrefix + "/role"

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


                dialogFormVisible: false,
                dialogTypeName: '修改',  //创建
                dialogType: 1,

                form: {
                    role: {},
                },

                rules: {
                    // title: [
                    //     {required: true, message: '不能为空', trigger: 'blur'},
                    //     {type: 'string', message: '长度需要在1~50之间', min: 1, max: 50}
                    // ],
                },

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

            //获取数据
            get_table_data() {
                this.get_method()
            },

            get_method(pagenum) {
                this.load_data = true
                // let this = this
                this.$axios.get(url)
                    .then((response) => {
                        this.load_data = false
                        if (response.data.success === true) {
                            this.table_data = response.data.result
                        }

                        global_.handleResponseGetThen(response)
                    })
                    .catch((error) => {
                        this.load_data = false
                        this.$message.error(error.response.data.message);
                    })
            },


            showPut(data) {
                this.dialogType = 2
                this.dialogTypeName = '修改'

                this.form = JSON.parse(JSON.stringify(data))

                this.dialogFormVisible = true


                // this.$axios.get(url + '/' + data.id)
                //     .then((response) => {
                //         if (response.data.success === true) {
                //             this.form = JSON.parse(JSON.stringify(response.data.result))
                //             this.dialogFormVisible = true
                //         }
                //     })
                //     .catch((error) => {
                //     })
            },

            postOrPut() {
                this.dialogFormVisible = false;

//                post
                if (this.dialogType === 1) {
                    // this.form.userIds1 = this.userIds1
                    // this.form.userIds2 = this.userIds2
                    var str = JSON.stringify(this.form)

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
                    // var form = JSON.parse(JSON.stringify(this.form))
                    // delete form.createTime
                    // delete form.updateTime

                    let form = {
                        roleName: this.form.role.roleName,
                        roleNameShow: this.form.role.roleNameShow,
                    }

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


        }
    }
</script>
