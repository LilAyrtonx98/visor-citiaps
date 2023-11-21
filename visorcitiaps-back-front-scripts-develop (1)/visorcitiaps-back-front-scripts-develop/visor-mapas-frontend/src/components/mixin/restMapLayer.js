export const restMapLayer = {
    data() {
        return {
            layer: {
              id_category: '-1',
              name: '',
              desc: '',
              geoserver_data: {
                  filename: '',
                  coordinates_system: 'EPSG:3857',
                  workspace: '',
                  datastore: ''
              },
            },
            layers: [],
        }
    },
    methods: {
      getLayer: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/layers/' + this.$route.params.idLayer)
        .then(function (response) {
            that.layer = response.data;
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getLayersPages: function(page){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/layers?page=' + page)
        .then(function (response) {
            that.layers = response.data;
            response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      postLayer: function(){
        var that = this;
        this.$http.post(this.$store.getters.backendurl + 'api/v1/layers', 
        this.layer
        )
        .then(function (response) {
            that.hideModalUpload();
            that.$emit('onMapToolAction', 'assignUsers','show', response.data.id)
        })
        .catch(function (error) {
            console.log(error);
        });
      },
      putLayer: function(){
        var that = this;
        this.$http.put(this.$store.getters.backendurl + 'api/v1/layers/' + this.$route.params.idLayer,
        this.layer
        )
        .then(function (response) {
            that.$router.push('/capas/'+ that.$route.params.idLayer)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      deleteLayer: function(){
        var that = this;
        this.$http.delete(this.$store.getters.backendurl + 'api/v1/layers/' + this.$route.params.idLayer)
        .then(function (response) {
            that.$router.push('/capas')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
    }
  }