import GeoJSON from 'ol/format/GeoJSON.js';
import VectorLayer from 'ol/layer/Vector.js';
import VectorSource from 'ol/source/Vector.js';
import {Fill, Stroke, Style, Text} from 'ol/style.js';

import {vuex} from '@/components/mixin/vuex.js'

var style = new Style({
  fill: new Fill({
    color: 'rgba(255, 255, 255, 0.6)'
  }),
  stroke: new Stroke({
    color: '#319FD3',
    width: 1
  }),
  text: new Text({
    font: '12px Calibri,sans-serif',
    fill: new Fill({
      color: '#000'
    }),
    stroke: new Stroke({
      color: '#fff',
      width: 3
    })
  })
});

var highlightStyle = new Style({
  stroke: new Stroke({
    color: '#f00',
    width: 1
  }),
  fill: new Fill({
    color: 'rgba(255,0,0,0.1)'
  }),
  text: new Text({
    font: '12px Calibri,sans-serif',
    fill: new Fill({
      color: '#000'
    }),
    stroke: new Stroke({
      color: '#f00',
      width: 3
    })
  })
});

var vectorLayerChile;

var featureOverlay;

var highlight;

var feature;

export const charts = {
    data() {
        return {
          layerType: ''
        }
    },
    mixins:[vuex],
    methods: {
      createChartVectorLayer: function (filename) {
        var that = this;
        that.selectMapControl('loadScreen', true)
        vectorLayerChile = new VectorLayer({
          renderMode: 'image',
          source: new VectorSource({
            url: 'data/' + filename + '.geojson',
            format: new GeoJSON()
          }),
          style: function(feature) {
            style.getText().setText(feature.get('NOM_' + that.layerType));
            return style;
          }
        });

        this.map.addLayer(vectorLayerChile);

        vectorLayerChile.on('change', function(e) {
          if (vectorLayerChile.getSource().getState() == 'ready') {
          }
          that.selectMapControl('loadScreen', false)
        });
        
        featureOverlay = new VectorLayer({
          renderMode: 'image',
          source: new VectorSource(),
          style: function(feature) {
            highlightStyle.getText().setText(/*feature.get(that.layerType)*/);
            return highlightStyle;
          }
        });

        this.map.addLayer(featureOverlay);
        
      },
      displayFeatureInfo: function (pixel) {

        feature = this.map.forEachFeatureAtPixel(pixel, function(feature) {
          return feature;
        });

        /*
        var info = document.getElementById('info');
        if (feature) {
          info.innerHTML = feature.getId() + ': ' + feature.get('name');
        } else {
          info.innerHTML = '&nbsp;';
        }
        */

        if (feature !== highlight) {
        if (highlight) {
          featureOverlay.getSource().removeFeature(highlight);
        }
        if (feature) {
          featureOverlay.getSource().addFeature(feature);
        }
        highlight = feature;
        }

      },

      // Realiza el procesamiento de los datos
      sendLayerDataForChart: function (pixel) {

        var that = this;

        feature = this.map.forEachFeatureAtPixel(pixel, function(feature) {
          return feature;
        });

        var data = [];

        var populationData = {
          isData: true,
          extra: {
            title: 'POBLACIÓN POR EDAD',
            subtitle: this.layerType,
            stats: [
              {number: 0, label: ' niños'}, 
              {number: 0, label: ' adultos'}, 
              {number: 0, label: ' adultos mayores'},
              {number: 0, label: ' en pobreza multidimensional'},
              //{number: 0, label: ' mujeres'}, 
              //{number: 0, label: ' hombres'},
            ]
          },
          series: [{
            name: 'Hombres',
            data: []
            },
            {
              name: 'Mujeres',
              data: []
            }
          ],
          chartOptions: {
            chart: {
              stacked: true,
              toolbar: {
                show: false
              }
            },
            colors: ['#008FFB', '#FF4560'],
            plotOptions: {
              bar: {
                horizontal: true,
                barHeight: '100%'
              },
            },
            dataLabels: {
              enabled: false
            },
            stroke: {
              width: 1,
              colors: ["#fff"]
            },
            grid: {
              xaxis: {
                showLines: false
              }
            },
            yaxis: {
              min: -1,
              max: 1,
              title: {
                //text: 'Años',
              }
            },
            tooltip: {
              shared: false,
              x: {
                formatter: function (val) {
                  return val + ' años';
                }
              },
              y: {
                formatter: function (val) {
                  return that.numberFormatter(val) + ' personas';
                }
              }
            },
            xaxis: {
              categories: [],
              title: {
                text: ' '
              },
              labels: {
                trim: false,
                formatter: function (val) {
                  return that.numberFormatter(Math.round(val));
                },
              }
            },
            plotOptions: {
              bar: {
                  horizontal: true,
                  endingShape: 'flat',
                  columnWidth: '100%',
                  barHeight: '100%',
                  distributed: false,
                  dataLabels: {
                      position: 'bottom',
                      maxItems: 100,
                      hideOverflowingLabels: false,
                  }
              }
            },
            legend: {
              position: 'top'
            }
          }
        }

        var fireData = {
          isData: true,
          extra: {
            title: 'INCENDIOS FORESTALES POR TEMPORADA',
            subtitle: this.layerType,
            stats: [
              {number: 0, label: ' incendios en promedio'}, 
              {number: 0, label: ' incendios en la última temporada'},
            ]
          },
          series: [{
            name: "Total",
            data: []
          }],
          chartOptions: {
            chart: {
              stacked: false,
              toolbar: {
                show: false
              },
              zoom: {
                enabled: false
              }
            },
            tooltip: {
              shared: false,
              x: {
                format: 'yyyy',
                formatter: function (val) {
                  let aux = new Date(val)
                  return 'Temporada ' + (Number(aux.getFullYear())-1) + '-' + aux.getFullYear();
                }
              },
              y: {
                formatter: function (val) {
                  return val + " incendios"
                }
              }
            },
            dataLabels: {
                enabled: false
            },
            stroke: {
                curve: 'straight'
            },
            colors: ['#00AC20'],
            xaxis: {
                type: 'datetime',
                categories: [],
                tickAmount: 6,           
            }
          }
        }

        var houseData = {
          isData: true,
          extra: {
            title: 'VIVIENDAS POR TIPO',
            subtitle: this.layerType,
            stats: [
              {number: 0, label: ' viviendas en total'},
              {number: 0, label: ' viviendas ocupadas'}, 
              {number: 0, label: ' viviendas desocupadas'}, 
              {number: 0, label: ' viviendas colectivas'}
            ]
          },
          series: [],
          chartOptions: {
            labels: [],
            legend: {
              position: 'bottom',
              horizontalAlign: 'center',
            },
            tooltip: {
              shared: false,
              y: {
                formatter: function (val) {
                  return that.numberFormatter(val) + ' viviendas';
                }
              }
            }
          }
        }

        var iH = 0;
        var iM = 0;
        var iF = 0;
        var iC = 0;

        if (feature) {
          for (let i = 0; i < feature.getKeys().length; i++) {
            //  Datos edad hombres
            if(feature.getKeys()[i].includes('HOMBRES_') && !feature.getKeys()[i].includes('Total') ) {

              iH = iH + 1;

              populationData.series[0].data.push(-1 * Number(feature.get(feature.getKeys()[i])));
              populationData.chartOptions.xaxis.categories.push(feature.getKeys()[i].replace("HOMBRES_", ""));

              if(populationData.chartOptions.yaxis.max < Number(feature.get(feature.getKeys()[i]))) {
                populationData.chartOptions.yaxis.max = Number(feature.get(feature.getKeys()[i]));
              }

              if(iH<=3){
                populationData.extra.stats[0].number = populationData.extra.stats[0].number + Number(feature.get(feature.getKeys()[i]));
              } else if(iH>=14) {
                populationData.extra.stats[2].number = populationData.extra.stats[2].number + Number(feature.get(feature.getKeys()[i]));
              } 
              else {
                populationData.extra.stats[1].number = populationData.extra.stats[1].number + Number(feature.get(feature.getKeys()[i]));
              }
              //populationData.extra.stats[5].number = populationData.extra.stats[4].number + Number(feature.get(feature.getKeys()[i]));
              
            }
            //  Datos edad mujeres
            if(feature.getKeys()[i].includes('MUJERES_') && !feature.getKeys()[i].includes('Total') ) {

              iM = iM + 1;

              populationData.series[1].data.push(Number(feature.get(feature.getKeys()[i])));

              if(populationData.chartOptions.yaxis.max < Number(feature.get(feature.getKeys()[i]))) {
                populationData.chartOptions.yaxis.max = Number(feature.get(feature.getKeys()[i]));
              }

              if(iM<=3){
                populationData.extra.stats[0].number = populationData.extra.stats[0].number + Number(feature.get(feature.getKeys()[i]));
              } else if(iM>=13) {
                populationData.extra.stats[2].number = populationData.extra.stats[2].number + Number(feature.get(feature.getKeys()[i]));
              }
              else {
                populationData.extra.stats[1].number = populationData.extra.stats[1].number + Number(feature.get(feature.getKeys()[i]));
              }
              //populationData.extra.stats[4].number = populationData.extra.stats[3].number + Number(feature.get(feature.getKeys()[i]));
              
            }
            if(feature.getKeys()[i].includes('POBREZA_') && !feature.getKeys()[i].includes('Total') ) {
              populationData.extra.stats[3].number = Number(feature.get(feature.getKeys()[i])).toString().replace(".", ",") + '%';
            }
            // Datos tipos de viviendas
            if(feature.getKeys()[i].includes('VIVIENDAS_') && !feature.getKeys()[i].includes('Total') ) {
              iC = iC + 1;

              houseData.series.push(Number(feature.get(feature.getKeys()[i])));
              houseData.chartOptions.labels.push(feature.getKeys()[i].replace("VIVIENDAS_", ""));

              houseData.extra.stats[0].number = houseData.extra.stats[0].number + Number(feature.get(feature.getKeys()[i]));
              
              if(iC<=2) {
                houseData.extra.stats[1].number = houseData.extra.stats[1].number + Number(feature.get(feature.getKeys()[i]));
              } else if(iC<=4) {
                houseData.extra.stats[2].number = houseData.extra.stats[2].number + Number(feature.get(feature.getKeys()[i]));
              } else {
                houseData.extra.stats[3].number = Number(feature.get(feature.getKeys()[i]));
              }
            }
            // Datos incendios por temporada (por región)
            if(feature.getKeys()[i].includes('CANTIDAD_INCENDIOS_REGIONALES_') && !feature.getKeys()[i].includes('Total') ) {
              iF = iF + 1;

              fireData.series[0].data.push(Number(feature.get(feature.getKeys()[i])));
              fireData.chartOptions.xaxis.categories.push(Number(feature.getKeys()[i].replace("CANTIDAD_INCENDIOS_REGIONALES_", "").replace(" - ","-").split("-")[1]) + '-12-31T00:00:00');
              
              fireData.extra.stats[0].number = fireData.extra.stats[0].number + Number(feature.get(feature.getKeys()[i]));
              fireData.extra.stats[1].number = Number(feature.get(feature.getKeys()[i]));
            }
          }
          populationData.chartOptions.yaxis.min = -populationData.chartOptions.yaxis.max;
          populationData.series[0].data = populationData.series[0].data.reverse();
          populationData.series[1].data = populationData.series[1].data.reverse();
          populationData.chartOptions.xaxis.categories = populationData.chartOptions.xaxis.categories.reverse();

          fireData.extra.stats[0].number = (fireData.extra.stats[0].number/iF).toFixed();

          populationData.extra.stats[0].number = this.numberFormatter(populationData.extra.stats[0].number);
          populationData.extra.stats[1].number = this.numberFormatter(populationData.extra.stats[1].number);
          populationData.extra.stats[2].number = this.numberFormatter(populationData.extra.stats[2].number);

          houseData.extra.stats[0].number = this.numberFormatter(houseData.extra.stats[0].number);
          houseData.extra.stats[1].number = this.numberFormatter(houseData.extra.stats[1].number);
          houseData.extra.stats[2].number = this.numberFormatter(houseData.extra.stats[2].number);
          houseData.extra.stats[3].number = this.numberFormatter(houseData.extra.stats[3].number);

          fireData.extra.stats[0].number = this.numberFormatter(fireData.extra.stats[0].number);
          fireData.extra.stats[1].number = this.numberFormatter(fireData.extra.stats[1].number);

          switch (this.layerType) {
            case "REGION":
              //populationData.extra.subtitle = this.firstCharCapital(feature.get('NOM_' + this.layerType));
              populationData.extra.subtitle = feature.get('NOM_' + this.layerType);
              houseData.extra.subtitle = feature.get('NOM_' + this.layerType);
              fireData.extra.subtitle = feature.get('NOM_' + this.layerType);
              break;
            case "COMUNA":
              //populationData.extra.subtitle = "Comuna de " + this.firstCharCapital(feature.get('NOM_' + this.layerType));
              populationData.extra.subtitle = "COMUNA DE " + feature.get('NOM_' + this.layerType);
              houseData.extra.subtitle = "COMUNA DE " + feature.get('NOM_' + this.layerType);
              fireData.extra.subtitle = "COMUNA DE " + feature.get('NOM_' + this.layerType) + " (DATOS REGIONALES)";
              break;
            default:
              break;
          }

        } else {
          populationData.isData = false;
          fireData.isData = false;
          houseData.isData = false;
        }

        if (feature !== highlight) {
        if (highlight) {
          featureOverlay.getSource().removeFeature(highlight);
        }
        if (feature) {
          featureOverlay.getSource().addFeature(feature);
        }
        highlight = feature;
        }

        data.push(populationData);
        data.push(fireData);
        data.push(houseData);

        this.executeMapToolAction('chart', 'showDataChart', {chartData: data})
      },
      chartClickCallback: function(evt) {
        this.sendLayerDataForChart(evt.pixel);
      },
      addOrRemoveChartLayerListener: function(chartLayer) {
        var that = this;
        
        switch (chartLayer) {
          case 'chileRegiones2018':
            this.layerType = 'REGION';
            this.removeChartLayer();
            this.createChartVectorLayer(chartLayer);
            this.map.addEventListener('click', this.chartClickCallback, false);
            break;
          case 'chileComunas2018':
            this.layerType = 'COMUNA';
            this.removeChartLayer();
            this.createChartVectorLayer(chartLayer);
            this.map.addEventListener('click', this.chartClickCallback, false);
            break;
          default:
            this.removeChartLayer();
            break;
        }

        /*
        this.map.on('pointermove', function(evt) {
          if (evt.dragging) {
            return;
          }
          var pixel = that.map.getEventPixel(evt.originalEvent);
          that.displayFeatureInfo(pixel);
        });
        */
      },
      removeChartLayer: function() {
        highlight = null;
        feature = null;
        this.map.removeLayer(vectorLayerChile);
        this.map.removeLayer(featureOverlay);
        this.map.removeEventListener('click', this.chartClickCallback, false);
        vectorLayerChile = null;
        featureOverlay = null;
      }
    }
  }