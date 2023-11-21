export const restGroup = {
    data() {
        return {
            group: {
              name: '',
              desc: '',
            },
            groups: []
        }
    },
    methods: {
      getGroup: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/' + this.$route.params.idGroup)
        .then(function (response) {
            that.group = response.data;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getGroupsPages: function(page){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/groups?page=' + page)
        .then(function (response) {
            that.groups = response.data;
            response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getGroupsNoPages: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/groups/')
        .then(function (response) {
            that.groups = response.data;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      putGroup: function(){
        var that = this;
        this.$http.put(this.$store.getters.backendurl + 'api/v1/groups/' + this.$route.params.idGroup,
        this.group
        )
        .then(function (response) {
            that.$router.push('/usuarios/grupos-de-usuarios/'+ that.$route.params.idGroup)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      deleteGroup: function(){
        var that = this;
        this.$http.delete(this.$store.getters.backendurl + 'api/v1/groups/' + this.$route.params.idGroup)
        .then(function (response) {
            that.$router.push('/usuarios/grupos-de-usuarios')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
    }
  }