export const restUser = {
    data() {
        return {
            user: {
                id_group: '-1',
                firstname: '',
                lastname: '',
                username: '',
                password: '',
                permissions: {
                    users: false,
                    layers: false,
                    geo: false,
                    maps: false,
                    visor: false,
                    annotation: false
                }
              },
              users: []
      }
    },
    methods: {
      getUser: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/users/' + this.$route.params.idUser)
        .then(function (response) {
            that.user = response.data;
            that.actualUsername = response.data.username;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getUsersPages: function(page){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/users?page=' + page)
        .then(function (response) {
            that.users = response.data;
            response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getUsersByGroup: function(page,idGroup){
        var that=this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/'+idGroup+'/users?page='+page)
        .then(function (response){
          that.users = response.data;
          response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
      },
      postUser: function(){
        var that = this;
        this.$http.post(this.$store.getters.backendurl + 'api/v1/users', 
        this.user
        )
        .then(function (response) {
          let idUser = response.data.id
          that.$router.push('/usuarios/'+ idUser + '/asignar-mapas')
        })
        .catch(function (error) {
            console.log(error);
        });
      },
      putUser: function(){
        var that = this;
        this.$http.put(this.$store.getters.backendurl + 'api/v1/users/' + this.$route.params.idUser,
        this.user
        )
        .then(function (response) {
            that.$router.push('/usuarios/'+ that.$route.params.idUser)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      deleteUser: function(){
        var that = this;
        this.$http.delete(this.$store.getters.backendurl + 'api/v1/users/' + this.$route.params.idUser)
        .then(function (response) {
            that.$router.push('/usuarios')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
    }
  }