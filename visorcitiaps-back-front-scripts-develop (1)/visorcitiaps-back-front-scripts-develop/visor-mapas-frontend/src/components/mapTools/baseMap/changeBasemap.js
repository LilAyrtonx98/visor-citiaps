import {Tile as TileLayer} from 'ol/layer.js';
import XYZ from 'ol/source/XYZ';
import {OSM} from 'ol/source.js';


var openStreetMap = new TileLayer({
  source: new OSM()
});

var local = new TileLayer({
  source: new XYZ({
    url: process.env.VUE_APP_MAPTILES
  })
})

export const changeBasemap = {
    data() {
        return {
            message: 'Hola'
        }
    },
    methods: {
      changeBasemap: function(basemap) {

        switch (basemap) {
          case 'local':
            openStreetMap.setVisible(false);
            local.setVisible(true);
            this.map.getView().setMaxZoom(12);
            break;
          default:
            local.setVisible(false);
            openStreetMap.setVisible(true);
            this.map.getView().setMaxZoom(18);
            break;
        }
      },
      getLocalBasemap: function() {
        return local
      },
      getOpenStreetMapBasemap: function() {
        return openStreetMap
      }

    }
  }