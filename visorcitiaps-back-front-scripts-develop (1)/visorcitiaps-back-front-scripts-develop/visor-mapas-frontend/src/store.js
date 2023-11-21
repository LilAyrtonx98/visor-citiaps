import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    role: undefined,
    token: null,
    mobile: false,

    // TRUE si se está mostrando un mapa, FALSE en caso contrario
    mapActive: false,

    // Almacena el estado de los controles de navegación y otros elementos del mapa
    mapControls: {
      legends: {active: false},
      loadScreen: {active: false},
      navDrawer: {active: false}
    },

    // Almacena la herramienta seleccionada del mapa
    mapTools: {
      select: {active: false, action: '', data: {}},
      measures: {active: false, action: '', data: {}},
      chart: {active: false, action: '', data: {}},
      magnify: {active: false, action: ''},
      export: {active: false, action: '', format: 'image', data: null},
      layers: {active: false, action: ''},
      search: {active: false, action: ''},
      geoprocessing: {active: false, action: ''},
      export: {active: false, action: ''},
      basemap: {active: false, action: ''},
      help: {active: false, action: ''},
      annotations: {active: false, action: '', data: {}},
      attributesTable: {active: false, action: '', data: {}}
    },

    layers: [],
    activeLayers: [],
    geoprocessings: [],

  },
  getters: {
    backendurl (){
      //Acceder de la siguiente forma:
      //this.$store.getters.backendurl
      return process.env.VUE_APP_API_URL;
    },
    getPermissions (state) {
      if(state.role){
        return state.role;
      }
      return undefined
    },
    getToken (state) {
      if(state.token){
        return state.token;
      }
      return undefined
    },
    // TRUE si se trata de un dispositivo móvil, FALSE si es un computador
    isMobile (state) {
      return state.mobile
    },
    isMapActive (state) {
      return state.mapActive
    },
    // Entrega la herramienta del mapa seleccionada 
    getMapSelectedTool: state => {
      let result = Object.keys(state.mapTools).filter(tool => state.mapTools[tool].active)
      return result.length > 0 ? result[0] : ''
    },
    // Entrega la herramienta que está ejecutando una acción
    getMapActionTool: state => {
      let result = Object.keys(state.mapTools).filter(tool => state.mapTools[tool].action != '')
      let result2 = result.length > 0 ? result[0] : ''
      return result2
    },
    // Entrega los atributos de una herramienta del mapa
    getMapTools: (state) => {
      return state.mapTools
    },
    // Entrega los atributos de un control de navegación del mapa
    getMapControls: (state) => {
      return state.mapControls
    },
    // Entrega las capas disponibles para el usuario del visor 
    getLayers: state => {
      return state.layers
    },
    // Entrega las capas seleccionadas para visualizarse en el mapa
    getActiveLayers: state => {
      return state.activeLayers
    },
    // Entrega los geoprocesos disponibles para el usuario del visor
    getGeoprocessings: state => {
      return state.geoprocessings
    },
  },
  mutations: {
    setPermissions (state){
      var name = "session_public" + "=";
      var decodedCookie = decodeURIComponent(document.cookie);
      var ca = decodedCookie.split(';');
      for(var i = 0; i <ca.length; i++) {
          var c = ca[i];
          while (c.charAt(0) == ' ') {
          c = c.substring(1);
          }
          if (c.indexOf(name) == 0) {
          state.role = JSON.parse(c.substring(name.length, c.length));
          return;
        }
      }
      state.role = undefined;
      return;
    },
    setToken (state){
      var name = "session_token" + "=";
      var decodedCookie = decodeURIComponent(document.cookie);
      var ca = decodedCookie.split(';');
      for(var i = 0; i <ca.length; i++) {
          var c = ca[i];
          while (c.charAt(0) == ' ') {
          c = c.substring(1);
          }
          if (c.indexOf(name) == 0) {
          state.token = c.substring(name.length, c.length);
          return;
        }
      }
      state.token = undefined;
      return;
    },
    setIsMobile (state){
      state.mobile = true;
      return
    },
    setIsNotMobile (state) {
      state.mobile = false;
      return
    },
    setMapActive (state, isActive) {
      state.mapActive = isActive
    },
    // Desactiva cualquier herramienta del mapa seleccionada anteriormente
    resetMapToolsSelection (state) {
      state.mapTools.select.active = false
      state.mapTools.measures.active = false
      state.mapTools.chart.active = false
      state.mapTools.magnify.active = false
      state.mapTools.layers.active = false
      state.mapTools.search.active = false
      state.mapTools.geoprocessing.active = false
      state.mapTools.export.active = false
      state.mapTools.basemap.active = false
      state.mapTools.help.active = false
      state.mapTools.annotations.active = false
      state.mapTools.attributesTable.active = false
    },
    // Selecciona una herramienta del mapa
    setMapToolSelection(state, payload) {
      state.mapTools[payload.tool].active = payload.active
    },
    // Elimina cualquier acción realizada por la herramienta
    resetMapToolsAction (state) {
      state.mapTools.select.action = ''
      state.mapTools.measures.action = ''
      state.mapTools.chart.action = ''
      state.mapTools.magnify.action = ''
      state.mapTools.layers.action = ''
      state.mapTools.search.action = ''
      state.mapTools.geoprocessing.action = ''
      state.mapTools.export.action = ''
      state.mapTools.basemap.action = ''
      state.mapTools.help.action = ''
      state.mapTools.annotations.action = ''
      state.mapTools.attributesTable.action = ''
    },
    // Ejecuta una acción en una herramienta del mapa
    setMapToolAction(state, payload) {
      state.mapTools[payload.tool].action = payload.action
      state.mapTools[payload.tool].data = payload.data
    },
    // Establece la lista de capas disponibles para el usuario
    setLayers(state, layers) {
      state.layers = layers
    },
    // Agrega una nueva capa a la lista de capas seleccionadas por el usuario para visualizarse en el mapa
    addActiveLayer(state, layer) {
      state.activeLayers.push(layer)
      console.log('add layer to map',layer)
    },
    // Remueve una capa de la lista de capas seleccionadas por el usuario para visualizarse en el mapa
    removeActiveLayer(state, layer) {
      let index = state.activeLayers.findIndex(function(item){ return item.id == layer.id})
      if (index > -1) {
        state.activeLayers.splice(index, 1);
      }
    },
    // Establece la lista de geoprocesos disponibles para el usuario
    setGeoprocessings(state, geoprocessings) {
      state.geoprocessings = geoprocessings
    },
    // Selecciona un control del mapa
    setMapControlSelection(state, payload) {
      state.mapControls[payload.control].active = payload.active
    },
    cleanActiveLayers(state) {
      state.activeLayers = []
    }
  },
  actions: {
    // Desactiva todas las herramientas de mapa
    resetMapToolsSelection ({ commit }) {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          commit('resetMapToolsSelection')
          resolve()
        }, 1) // Delay mínimo para asegurar que la siguiente acción se ejecute después de esta
      })
    },
    // Selecciona una herramienta en el mapa
    setMapToolSelection ({ dispatch, commit }, payload) {
      if(payload.tool=='attributesTable'){
        commit('setMapToolSelection', payload)
      }
      else{
        return dispatch('resetMapToolsSelection').then(() => {
          commit('setMapToolSelection', payload)
        })
      }
    },
    // Elimina todas las acciones ejecutadas anteriormente
    resetMapToolsAction ({ commit }) {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          commit('resetMapToolsAction')
          resolve()
        }, 1) // Delay mínimo para asegurar que la siguiente acción se ejecute después de esta
      })
    },
    // Ejecuta una acción en una herramienta
    setMapToolAction ({ dispatch, commit }, payload) {
      return dispatch('resetMapToolsAction').then(() => {
        commit('setMapToolAction', payload)
      })
    },
    // Establece la lista de capas disponibles para el usuario
    setLayers(context, layers) {
      context.commit('setLayers', layers)
    },
    // Agrega una nueva capa a la lista de capas seleccionadas por el usuario
    addActiveLayer(context, layer) {
      context.commit('addActiveLayer', layer)
    },
    // Elimina una capa de la lista de capas seleccionadas por el usuario
    removeActiveLayer(context, layer) {
      context.commit('removeActiveLayer', layer)
    },
    // Establece la lista de geoprocesos disponibles para el usuario
    setGeoprocessings(context, geoprocessings) {
      context.commit('setGeoprocessings', geoprocessings)
    },
    // Selecciona un control del mapa
    setMapControlSelection(context, payload) {
      context.commit('setMapControlSelection', payload)
    },
    // Reiniciar el estado del mapa
    resetMap(context) {
      context.commit('resetMapToolsSelection')
      context.commit('resetMapToolsAction')
    },
    // Selecciona un control del mapa
    cleanActiveLayers(context) {
      context.commit('cleanActiveLayers')
    },
  }
})
