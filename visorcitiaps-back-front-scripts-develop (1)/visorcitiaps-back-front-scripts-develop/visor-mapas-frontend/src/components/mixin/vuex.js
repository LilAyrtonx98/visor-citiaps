
export const vuex = {
    data() {
        return {
        }
    },
    computed: {
      // TRUE si es un dispositivo móvil, FALSE si es un computador
      isMobile: function() {
        return this.$store.getters.isMobile
      },
      // TRUE si se está mostrando un mapa, FALSE en caso contrario
      isMapActive: function() {
        return this.$store.getters.isMapActive
      },
      // Entrega la herramienta del mapa seleccionada
      mapToolSelected: function() {
        return this.$store.getters.getMapSelectedTool
      },
      // Entrega la herramienta del mapa que ha ejecutado una acción
      mapToolExecuted: function() {
        return this.$store.getters.getMapActionTool
      },
      // Entrega la lista de capas disponibles para el usuario
      layers: function() {
        return this.$store.getters.getLayers
      },
      // Entrega la lista de capas seleccionadas por el usuario para visualizarse en el mapa
      activeLayers: function() {
        return this.$store.getters.getActiveLayers
      },
      // Entrega la lista de geoprocesos disponiles para el usuario
      geoprocessings: function() {
        return this.$store.getters.getGeoprocessings
      }
    },
    methods: {
      // Entrega los atributos de una herramienta del mapa
      mapToolValues: function(toolname) {
        return this.$store.getters.getMapTools[toolname]
      },
      // Entrega los atributos de un control del mapa
      mapControlValues: function(controlname) {
        return this.$store.getters.getMapControls[controlname]
      },
      // Selecciona una herramienta del mapa
      selectMapTool: function(toolname, isActive) {
        this.$store.dispatch('setMapToolSelection', {
            tool: toolname,
            active: isActive
        })
      },
      // Selecciona un control del mapa
      selectMapControl: function(controlname, isActive) {
        this.$store.dispatch('setMapControlSelection', {
            control: controlname,
            active: isActive
        })
      },
      // Ejecuta una acción en una herramienta
      // toolname (String '') nombre de la herramienta
      // action (String '') nombre de la acción ejecutada
      // data (Object {}) atributos necesarios para realizar la acción
      executeMapToolAction: function(toolname, action, data) {
        this.$store.dispatch('setMapToolAction', {
            tool: toolname,
            action: action,
            data: data
        })
      },
      resetMap: function() {
        this.$store.dispatch('resetMap')
      },
      setMapActive: function(isActive) {
        this.$store.commit('setMapActive', isActive)
      },
      cleanActiveLayers: function() {
        this.$store.dispatch('cleanActiveLayers')
      }
    }
  }