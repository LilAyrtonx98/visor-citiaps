export const formatter = {
    data() {
        return {
          moment: require('moment')
        }
    },
    methods: {
      dateFormatter: function (date) {
        this.moment.locale('es');
        return this.moment(date).format('LL');
      },
      cleanString: function(cadena){
          // Definimos los caracteres que queremos eliminar
          var specialChars = "!@#$^&%*()+=-[]\/{}|:<>?,.`'\"";

          // Los eliminamos todos
          for (var i = 0; i < specialChars.length; i++) {
              cadena= cadena.replace(new RegExp("\\" + specialChars[i], 'gi'), '');
          }   

          // Lo queremos devolver limpio en minusculas
          cadena = cadena.toLowerCase();

          // Quitamos espacios y los sustituimos por _ porque nos gusta mas asi
          cadena = cadena.replace(/ /g,"_");

          // Quitamos acentos y "ñ". Fijate en que va sin comillas el primer parametro
          cadena = cadena.replace(/á/gi,"a");
          cadena = cadena.replace(/é/gi,"e");
          cadena = cadena.replace(/í/gi,"i");
          cadena = cadena.replace(/ó/gi,"o");
          cadena = cadena.replace(/ú/gi,"u");
          cadena = cadena.replace(/ä/gi,"a");
          cadena = cadena.replace(/ë/gi,"e");
          cadena = cadena.replace(/ï/gi,"i");
          cadena = cadena.replace(/ö/gi,"o");
          cadena = cadena.replace(/ü/gi,"u");
          cadena = cadena.replace(/ñ/gi,"n");
          return cadena;
        },
        numberFormatter: function(val) {
          return Math.abs(val).toString().replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1.');
        },
        firstCharCapital: function (string) {
          let result = '';
          let array = string.split(" ");
          for (let i = 0; i < array.length; i++) {
            array[i] = array[i].toLowerCase();
            if(array[i].length>3) {
              array[i] = array[i].charAt(0).toUpperCase() + array[i].slice(1);
            }
            result = result + array[i] + ' '; 
          }
          console.log(array)
          return result;
        }
    }
  }