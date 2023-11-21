import {OverviewMap, ScaleLine} from 'ol/control.js';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';

var source = new OSM();
var overviewMapControl = new OverviewMap({
  layers: [
    new TileLayer({
      source: source
    })
  ],
  className: 'ol-overviewmap ol-custom-overviewmap',
  collapsed: false,
  collapsible: false
});

export const mapControlsOpenLayers = {
    data() {
        return {

        }
    },
    methods: {
      addOverview: function() {
        this.map.addControl(overviewMapControl);
      },
      addScaleline: function() {
        var scaleline = new ScaleLine({
        })
        this.map.addControl(scaleline);
      },
      addControlsOpenLayers: function() {
        this.addOverview();
        this.addScaleline();
      },
    }
  }