export const restMapCategory = {
    data() {
        return {
            category: {
              name: '',
              desc: '',
            },
            categories: []
        }
    },
    methods: {
      getCategory: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/categories/' + this.$route.params.idCategory)
        .then(function (response) {
            that.category = response.data;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getCategoriesPages: function(page){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/categories?page=' + page)
        .then(function (response) {
            that.categories = response.data;
            response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getCategoriesNoPages: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/categories/')
        .then(function (response) {
            that.categories = response.data;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      postCategory: function(){
        var that = this;
        this.$http.post(this.$store.getters.backendurl + 'api/v1/categories', 
        this.category
        )
        .then(function (response) {
            //Obtener ID de la tarea recién creada y pasar a la vista de asignar usuarios
            let idCategory= response.data.id
            that.$router.push('/capas/categorias-de-capas/'+ idCategory)
        })
        .catch(function (error) {
            console.log(error);
        });
      },
      putCategory: function(){
        var that = this;
        this.$http.put(this.$store.getters.backendurl + 'api/v1/categories/' + this.$route.params.idCategory,
        this.category
        )
        .then(function (response) {
            that.$router.push('/capas/categorias-de-capas/'+ that.$route.params.idCategory)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      deleteCategory: function(){
        var that = this;
        this.$http.delete(this.$store.getters.backendurl + 'api/v1/categories/' + this.$route.params.idCategory)
        .then(function (response) {
            that.$router.push('/capas/categorias-de-capas/')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
    }
  }