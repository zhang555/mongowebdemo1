<script type="">
    import Vue from 'vue'

    const isDevelopment = process.env.NODE_ENV === 'development'


    // let url = window.location.href
    // let index = url.lastIndexOf("#")
    // let suffix = url.slice( 0 ,index - 1)
    let suffix = document.domain


    export default {

       // urlPrefix: isDevelopment ? "http://localhost:8228" : "http://192.168.0.112:8228",
        urlPrefix: isDevelopment ? "http://localhost:8228" : "http://"+suffix+":8228",

        pageSizeNum: 10,
        pageSizeString: "10",


        //
        handleResponseGetThen(response) {
            if (response.data.success === false) {
                if (response.data.code === 4011) {
                    Vue.prototype.$message("登陆过期，请重新登录")
                } else {
                    Vue.prototype.$message.error(response.data.message)
                }
                // Vue.prototype.$message("请重新登录")
            }
        },


        //
        handleResponseThen(response) {
            if (response.data.success === true) {
                Vue.prototype.$message.success(response.data.message)
            }
            if (response.data.success === false) {
                Vue.prototype.$message.error(response.data.message)
            }
        },


        //
        handleResponseThenCustomizePost(response) {
            if (response.data.success === true) {
                Vue.prototype.$message.success("创建成功")
            }
            if (response.data.success === false) {
                Vue.prototype.$message.error("创建失败")
            }
        },


        //
        handleResponseThenCustomizePut(response) {
            if (response.data.success === true) {
                Vue.prototype.$message.success("修改成功")
            }
            if (response.data.success === false) {
                Vue.prototype.$message.error("修改失败")
            }
        },

        //
        handleResponseCatch(error) {
            Vue.prototype.$message("失败了哦");
        },

        //
        handleUploadError(err) {
            Vue.prototype.$message("上传失败，请重试");
        },

    }
</script>


<style scoped>

</style>

