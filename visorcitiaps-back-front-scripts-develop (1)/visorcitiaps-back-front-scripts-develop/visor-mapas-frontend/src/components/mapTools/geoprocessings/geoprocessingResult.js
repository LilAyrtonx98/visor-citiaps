export const geoprocessingResult = {
  data() {
    return {
      message: 'Hola',
      data: {},
      dataChart: {},
      layersResult: [],
    }
  },
  methods: {
    receiveAttributes: function (attributes) {
      this.dataProcessingChart(attributes)
    },
    dataProcessingAverage: function (attributes) {
      // Nombre de capas
      var layersResult = []
      let keys = Object.keys(attributes[0])
      this.getLayersOfKeysAttributes(layersResult, keys)

      // Objetos por capa
      layersResult.forEach(layer => {
        this.data[layer] = {}
      })

      attributes.forEach(polygon => {
        let keys = Object.keys(polygon)
        keys.forEach(key => {

          if (!(key.includes('cut_') || key.includes('id_'))) {

            // El valor del polígono es un número?
            if (typeof polygon[key] == 'number') {
              let aux = key.split('.')
              // El valor del dato final está definido?
              if (isNaN(this.data[aux[0]][aux[1]])) {
                // Inicializar valores numéricos
                this.data[aux[0]][aux[1]] = 0
                this.data[aux[0]][aux[1]] = this.data[aux[0]][aux[1]] + polygon[key] / attributes.length
              } else {
                // Sumar valores
                this.data[aux[0]][aux[1]] = this.data[aux[0]][aux[1]] + polygon[key] / attributes.length
              }
            }

          }
        })
      });
    },
    dataProcessingChart: function (attributes) {
      // Nombre de capas
      var layersResult = []
      let keys = Object.keys(attributes[0])
      this.getLayersOfKeysAttributes(layersResult, keys)

      // Objetos por capa
      layersResult.forEach(layer => {
        this.dataChart[layer] = {}
      })

      attributes.forEach(polygon => {
        let keys = Object.keys(polygon)
        var nombresCapas = {}

        // Nombre polígonos
        keys.forEach(key => {
          if (!(key.includes('cut_') || key.includes('id_'))) {
            let aux = key.split('.')
            // El valor del polígono es un número?
            // Es un nombre?
            if (typeof polygon[key] == 'string') {
              if (key.includes('nombre') || key.includes('name')) {
                nombresCapas[aux[0]] = polygon[key]
              }
            }
          }
        })

        // Valores polígono
        keys.forEach(key => {

          if (!(key.includes('cut_') || key.includes('id_'))) {
            let aux = key.split('.')
            // El valor del polígono es un número?
            if (typeof polygon[key] == 'number') {
              // El valor del dato final es array?
              if (!Array.isArray(this.dataChart[aux[0]][aux[1]])) {
                // Inicializar arreglos numéricos
                this.dataChart[aux[0]][aux[1]] = []
                this.dataChart[aux[0]][aux[1]].push(
                  {
                    nombre: nombresCapas[aux[0]],
                    valor: polygon[key]
                  }
                )
              } else {
                // Agregar valores
                this.dataChart[aux[0]][aux[1]].push(
                  {
                    nombre: nombresCapas[aux[0]],
                    valor: polygon[key]
                  }
                )
              }
            }
          }
        })
      });
      this.layersResult = layersResult
    },
    getLayersOfKeysAttributes: function (layersResult, keys) {
      keys.forEach(key => {
        let layerName = key.split('.')[0]
        layersResult.includes(layerName) ? true : layersResult.push(layerName)
      })
    },
    transformDataForChart: function (layerData) {
      var barData = {
        isData: true,
        series: [

        ],
        chartOptions: {
          plotOptions: {
            bar: {
              distributed: true,
              horizontal: false,
              columnWidth: "75%",
              endingShape: "flat",
            }
          },
          colors: ['#036016'],
          dataLabels: {
            enabled: false
          },
          stroke: {
            show: true,
            width: 2,
            colors: ["transparent"]
          },

          xaxis: {
            title: {
              text: ""
            },
            categories: [

            ]
          },
          yaxis: {
            title: {
              text: ""
            },
            labels: {
              formatter: function (y) {
                return y.toFixed(2);
              }
            }
          },
          fill: {
            opacity: 1,
            
          },
          tooltip: {
            y: {
              formatter: function (val) {
                return  val;
              }
            }
          }
        }
      };

      let keys = Object.keys(layerData)
      keys.forEach(key => {
        let body = {}
        body.name = key
        body.data = []
        let categories = []

        layerData[key].forEach(datos => {
          body.data.push(datos.valor)
          categories.push(datos.nombre)
        })
        barData.series.push(body)
        barData.chartOptions.xaxis.categories = categories
      })
      return barData
    }
  }
}