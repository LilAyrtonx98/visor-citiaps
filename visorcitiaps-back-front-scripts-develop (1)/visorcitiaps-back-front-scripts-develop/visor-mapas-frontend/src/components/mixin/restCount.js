export const restCount = {
    data() {
        return {
          mapsCount: 0,  
          layersCount: 0,
          geoCount: 0,
          usersCount: 0,
          categoriesCount: 0,
          groupsCount: 0
        }
    },
    methods: {
      getMapsCount: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/counts/' + 'maps')
        .then(function (response) {
            that.mapsCount = response.data;
        })
        .catch(function (error) {
            console.log(error)
        });
      },
      getLayersCount: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/counts/' + 'layers')
        .then(function (response) {
            that.layersCount = response.data;
        })
        .catch(function (error) {
            console.log(error)
        });
      },
      getGeoCount: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/counts/' + 'geoprocessings')
        .then(function (response) {
            that.geoCount = response.data;
        })
        .catch(function (error) {
            console.log(error)
        });
      },
      getUsersCount: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/counts/' + 'users')
        .then(function (response) {
            that.usersCount = response.data;
        })
        .catch(function (error) {
            console.log(error)
        });
      },
      getCategoriesCount: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/counts/' + 'categories')
        .then(function (response) {
            that.categoriesCount = response.data;
        })
        .catch(function (error) {
            console.log(error)
        });
      },
      getGroupsCount: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/counts/' + 'groups')
        .then(function (response) {
            that.groupsCount = response.data;
        })
        .catch(function (error) {
            console.log(error)
        });
      },
    }
  }