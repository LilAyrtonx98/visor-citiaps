import Draw from 'ol/interaction/Draw.js';
import {Vector as VectorLayer} from 'ol/layer.js';
import {Vector as VectorSource} from 'ol/source.js';

//var typeSelect = document.getElementById('type');

var draw; // global so we can remove it later

var source = new VectorSource({wrapX: false});

var vector;

var coords;

var radius;


/**
 * Handle change event.
 */
/*
typeSelect.onchange = function() {
  map.removeInteraction(draw);
  addInteraction();
};
*/

source.on('addfeature', function(evt){
  var feature = evt.feature;

  if(feature.getGeometry().getCoordinates() == null) {
    // Si la figura está dibujada y no tiene coordenadas, es porque es un círculo y solo tiene un punto como centro

    coords = [feature.getGeometry().getCenter()]
    radius = feature.getGeometry().getRadius()
    console.log(feature.getGeometry().getRadius())
  } else {
    // Si la figura tiene coordenadas, es un polígono dibujado
    coords = feature.getGeometry().getCoordinates()
  }
  console.log(coords)
});

export const geoprocessingsInteraction = {
    data() {
        return {
            message: 'Hola',
        }
    },
    methods: {
      addGeoInteraction: function(drawType) {
        vector = new VectorLayer({
          source: source
        });
        this.map.addLayer(vector);
        draw = new Draw({
          source: source,
          type: drawType
          //type: 'Polygon'
        });
        draw.on('drawstart', function(evt){
          source.clear() // Elimina las capas anteriores
        })
        this.map.addInteraction(draw);
      },
      removeGeoInteraction: function() {
        this.map.removeLayer(vector);
        this.map.removeInteraction(draw);
        source.clear();
        draw = null;
        vector = null;
        coords = null;
      },
      cleanGeoDraw: function() {
        source.clear();
        coords = null;
      },
      getPolygonCoords: function() {
        return coords;
      },
      getCircleRadius: function() {
        return radius
      }
    },
    mounted () {
      
    }
  }