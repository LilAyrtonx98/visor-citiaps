<template>
  <div v-show="isActive" :class="{widgetContainer : !isFixed}" :id="isFixed? 'searchIDFix' : 'searchID'">
    <div v-if="!isFixed" :class="{widgetHeader : !isFixed}" id="searchIDheader">
        <p>Buscador</p>
        <div>
          <button class="btn btn-info" v-on:click="isHidden = !isHidden">
            <font-awesome-icon v-if="!isHidden" icon="minus"></font-awesome-icon>
            <font-awesome-icon v-else icon="plus"></font-awesome-icon>
          </button>
          <button class="btn btn-danger">
            <font-awesome-icon icon="times" v-on:click="exitSearch()"></font-awesome-icon>
          </button>
        </div>
    </div>

    <div :class="{widgetBody : !isFixed}" style="max-width: min-content;" v-show="!isHidden">
      <div :class="{resizable : !isFixed}">
        <div>
          <AutoCompleteInput
          ref="autoCompleteInput"
          id="location"
          v-bind:value="searchInput"
          v-bind:options="searchList"
          @updated="waitForSearch"
          @itemSelected="setLocation"
          v-bind:attribute="'display_name'"
          v-bind:beforeSearchMessage="'Ingresa la ubicación que estás buscando'"
          v-bind:searchingMessage="'Cargando resultados'"
          v-bind:emptySearchResultsMessage="'No hay resultados para esta ubicación'"
          v-bind:error="searchError"
          >
          </AutoCompleteInput>

          <div class="place-info">
            <div v-if="searchLocation">
              <div class="place-title">
                <img src="images/marker.svg">
                <p>Resultados búsqueda</p>
              </div>
              <p class="place-name">{{searchLocation.display_name}}</p>
              <p>
                <span class="place-lat-lon">
                  Latitud:
                </span>
                <span>
                  {{searchLocation.lat}}
                </span>
              </p>
              <p>
                <span class="place-lat-lon">
                  Longitud:
                </span>
                <span>
                  {{searchLocation.lon}}
                </span>
              </p>
            </div>

            <div v-else>
              <div class="place-title">
                <img src="images/marker.svg">
                <p>Resultados búsqueda</p>
              </div>
              <p class="place-name">No hay resultados para mostrar</p>
            </div>
          </div>
        </div>

        <div>
          <button class="btn btn-secondary btn-clean" style="margin-top: 0.5em;" v-if="searchLocation" v-on:click="cleanSearch()">Limpiar búsqueda</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import AutoCompleteInput from '@/components/mapTools/search/AutoCompleteInput.vue'
import {draggableDiv} from '@/components/mixin/draggableDiv.js'
import {vuex} from '@/components/mixin/vuex.js'


export default {
  name: 'HelloWorld',
  props: ['isFixed'],
  mixins:[vuex, draggableDiv],
  data() {
    return {
      isHidden: false,
      searchInput: '',
      searchList: [],
      searchLocation: '',
      timeout: null,
      searchError: false,
    }
  },
  components:{
    AutoCompleteInput
  },
  computed: {
    isActive: function() {
      return this.isFixed? this.isMobile : this.mapToolValues('search').active && !this.isMobile
    }
  },
  methods:{
    searchQuery: function(query){
      this.searchError = false;

      var req = new XMLHttpRequest();
      req.open('GET', process.env.VUE_APP_API_NOMINATIM + 'search?format=json&country=CL&q=' + query, false);
      req.withCredentials = false;
      req.send();

      if (req.status == 200){
        console.log("Success search query");
        console.log(req.response);
        if(req.response.lenght==0){
            this.searchError = true;
            return;
        }else{
            this.searchList = JSON.parse(req.response);
        }
      }else{
        console.log("Fail search query");
      }
        
      // DEPRECATED CORS ERROR needed withCredentials = false
      // this.$http.get( process.env.VUE_APP_API_NOMINATIM + 'search?format=json&country=CL&q=' + query)
      // .then(response => {
      //     if(response.data.lenght==0)
      //     {
      //         this.searchError = true;
      //         return;
      //     }
      //     this.searchList = response.data;
      // })
      // .catch((error) => {
      //     //Si se usa la api de locationiq y no hay resultados, la respuesta tiene un código de error (404)
      //     //Si falla la conuslta en cualquier caso se marca como error igual
      //     this.searchError = true;
      // })
    },
    waitForSearch: function (newValue) {
        this.searchInput = newValue;
    },
    setLocation: function(location)
    {
        this.searchLocation = location;
        this.executeMapToolAction('search', 'setLocation', {
          lat: location.lat,
          lon: location.lon
        })
    },
    exitSearch: function() {
      this.cleanSearch()
      this.selectMapTool('search', false)
    },
    cleanSearch: function() {
      this.$refs.autoCompleteInput.setQuery('')
      this.executeMapToolAction('search', 'removeLocation', {})
    }
  },
  updated () {

  },
  created () {

  },
  mounted () {
    this.makeDraggable("searchID");
  },
  watch: {
      'searchInput': function (newValue, oldValue) {

          if(newValue=='')
          {
            this.searchList = [];
            this.searchLocation = null;
            this.$emit('onRemoveLocation');
            return;
          }

          clearTimeout(this.timeout);
          // Make a new timeout set to go off in 800ms
          this.timeout = setTimeout(function () {
              this.searchQuery(newValue);
          }.bind(this), 501);
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.resizable{
  margin-top: 0.5em;
  height: 30vh;
  min-height: 10vh;
  min-width: 20vw;
  resize: both;
  overflow: auto;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.place-info {
  
  .place-title {
    display: flex;
    align-items: center;
    font-weight: bold;
    font-size: 1.25em;
    img {
      width: 1.25em;
      height: 1.25em;
    }
  }
  .place-name {
    //font-size: 1.0625em;
    //color: #036016;
  }
  .place-lat-lon{
    color: #036016;
  }
  span{
    //font-weight: bold;
  }
}

</style>
