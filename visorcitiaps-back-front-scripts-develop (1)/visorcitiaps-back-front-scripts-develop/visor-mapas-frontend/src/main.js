import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vClickOutside from 'v-click-outside'
 
Vue.use(vClickOutside)

//Axios
const axios = require('axios');
axios.defaults.withCredentials = true

// Respuesta global para axios si http code = 403 (backend lo entrega cuando no se encuentra la sesión)
const errorHandler = (error) => {
  if(error.response.request.status == 403) {
    console.log('Sesión no encontrada')
    // Eliminar cookies públicas (cierra sesión)
    window.location.reload()
    document.cookie = "session_public" + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    $store.commit('setPermissions').bind(this);
    $router.push('/inicio').bind(this)
  }
  return Promise.reject({ ...error })
}
const successHandler = (response) => {
  return response
}
axios.interceptors.response.use(
  response => successHandler(response),
  error => errorHandler(error)
)

Vue.prototype.$http = axios

//Validador global (español)
import VeeValidate, { Validator } from 'vee-validate';
import es from 'vee-validate/dist/locale/es'
Validator.localize({ es: es })
Vue.use(VeeValidate, {locale: 'es'})

//FontAwesome icons
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon, FontAwesomeLayers, FontAwesomeLayersText } from '@fortawesome/vue-fontawesome'

//Shortkeys
Vue.use(require('vue-shortkey'))

import VueHtml2Canvas from 'vue-html2canvas';
Vue.use(VueHtml2Canvas);


import { 
  faChevronDown, 
  faChevronUp, 
  faAngleDoubleLeft, 
  faAngleLeft, 
  faAngleRight, 
  faAngleDoubleRight, 
  faLayerGroup,
  faFileAlt,
  faQuestion,
  faToolbox,
  faFileImport,
  faRulerCombined,
  faBinoculars,
  faFileExport,
  faRuler,
  faMousePointer,
  faHandPointer,
  faChartBar,
  faSearch,
  faTimes,
  faMinus,
  faPlus,
  faMap,
  faChevronLeft,
  faArrowLeft,
  faChevronRight,
  faArrowUp,
  faArrowDown,
  faCogs,
  faCheck,
  faCopy,
  faStickyNote,
  faTrashAlt,
  faEllipsisH,
  faTable
  } from '@fortawesome/free-solid-svg-icons'

library.add( 
  faChevronDown, 
  faChevronUp, 
  faChevronLeft,
  faChevronRight,
  faArrowLeft,
  faAngleDoubleLeft, 
  faAngleLeft, 
  faAngleRight, 
  faAngleDoubleRight, 
  faLayerGroup,
  faFileAlt,
  faQuestion,
  faToolbox,
  faFileImport,
  faRulerCombined,
  faBinoculars,
  faFileExport,
  faRuler,
  faMousePointer,
  faHandPointer,
  faChartBar,
  faSearch,
  faTimes,
  faMinus,
  faPlus,
  faMap,
  faArrowUp,
  faArrowDown,
  faCogs,
  faCheck,
  faCopy,
  faStickyNote,
  faTrashAlt,
  faEllipsisH,
  faTable
  )
  
Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.component('font-awesome-layers', FontAwesomeLayers)
Vue.component('font-awesome-layers-text', FontAwesomeLayersText)

//Clipboard
import VueClipboard from 'vue-clipboard2'
Vue.use(VueClipboard);

//Modal
import VModal from 'vue-js-modal';
Vue.use(VModal,{dialog:true,dynamic:true});

//Gráficos
import VueApexCharts from 'vue-apexcharts'
Vue.use(VueApexCharts)
Vue.component('apexchart', VueApexCharts)

//Estilos (s)css
require("./assets/style/style.scss");
require("./assets/style/style-error.scss");
require("./assets/style/style-modal.scss");

// Estilos mapas
require("./assets/style/mapStyle/map.scss");
require("./assets/style/mapStyle/map-controls.css");
require("./assets/style/mapStyle/map-controls-icons.css");
require("./assets/style/mapStyle/map-tools.css");
require("./assets/style/mapStyle/map-popover.css");
//require("./assets/style/fontawesome/css/all.css"); //Producción

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
