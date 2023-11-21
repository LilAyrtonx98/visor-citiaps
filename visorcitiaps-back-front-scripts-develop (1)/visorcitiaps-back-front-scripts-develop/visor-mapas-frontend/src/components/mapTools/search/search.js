import {Vector as VectorLayer} from 'ol/layer.js';import XYZ from 'ol/source/XYZ';
import {Vector as VectorSource} from 'ol/source.js';
import {fromLonLat} from 'ol/proj.js';
import Feature from 'ol/Feature.js';
import Point from 'ol/geom/Point.js';
import {Style, Icon} from 'ol/style.js';

export const search = {
    data() {
        return {
            marker: null
        }
    },
    methods: {

      addLocationMarker: function(lat, lon){

        if(this.marker) {
          this.map.removeLayer(this.marker);
        }
        
        var pos = fromLonLat([parseFloat(lon), parseFloat(lat)]);
  
        var iconFeature = new Feature({
          geometry: new Point(pos),
        });
  
       var style = new Style({
          image: new Icon(/** @type {module:ol/style/Icon~Options} */ ({
            anchor: [0.5, 0.5],
            scale: 2,
            anchorXUnits: 'fraction',
            anchorYUnits: 'fraction',
            src: 'images/marker.svg'
          }))
        });
  
        iconFeature.setStyle(style)
  
        var vectorLayer = new VectorLayer({
          source: new VectorSource({
            features: [iconFeature]
          })
        });
  
        this.marker = vectorLayer;
  
        this.map.addLayer(this.marker)
  
        this.map.getView().animate({
            center: pos,
            duration: 1000,
            zoom: 15,
        });
      
      },
      removeLoationMarker: function() {
        this.map.removeLayer(this.marker);
        this.marker = null;
      },

    }
  }