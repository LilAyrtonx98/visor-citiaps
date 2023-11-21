export const restMap = {
    data() {
        return {
          map: {
            name: '',
            desc: '',
            imgurl: ''
          },
          maps: []
        }
    },
    methods: {
      getMap: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap)
        .then(function (response) {
            that.map = response.data;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getMapsPages: function(page){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/maps?page=' + page)
        .then(function (response) {
            that.maps = response.data;
            response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getMapsNoPages: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/')
        .then(function (response) {
            that.maps = response.data;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      postMap: function(){
        var that = this;
        this.$http.post(this.$store.getters.backendurl + 'api/v1/maps', 
        this.map
        )
        .then(function (response) {
          let idMap= response.data.id
          that.$router.push('/mapas/'+ idMap)
        })
        .catch(function (error) {
            console.log(error);
        });
      },
      putMap: function(){
        var that = this;
        this.$http.put(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap,
        this.map
        )
        .then(function (response) {
          that.$router.push('/mapas/'+ that.$route.params.idMap)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      deleteMap: function(){
        var that = this;
        this.$http.delete(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap)
        .then(function (response) {
            that.$router.push('/mapas')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getLayersByMap: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/layers/categories')
        .then(function (response) {
            console.log("HOLAAAAAAAA",response)
            that.$store.dispatch('setLayers', response.data)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getGeoprocessingsByMap: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/maps/' + this.$route.params.idMap + '/geoprocessings')
        .then(function (response) {
            that.$store.dispatch('setGeoprocessings', response.data)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
    }
  }