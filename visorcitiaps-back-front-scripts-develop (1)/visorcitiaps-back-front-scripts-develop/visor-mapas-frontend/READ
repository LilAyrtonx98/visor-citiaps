# visor-mapas-frontend

# PRIMERO LEER README REPO SCRIPTS (.env en raiz necesarios .env.production y .env.development)

## Visor de mapas y panel de administración

### Herramientas utilizadas
- [Vue.js](https://github.com/vuejs/vue) ([MIT License](https://github.com/vuejs/vue/blob/dev/LICENSE))
- [OpenLayers](https://github.com/openlayers/openlayers) ([2-Clause BSD License](https://github.com/openlayers/openlayers/blob/master/LICENSE.md))
- [ApexCharts.js (Vue wrapper)](https://github.com/apexcharts/vue-apexcharts) ([MIT License](https://github.com/apexcharts/vue-apexcharts/blob/master/LICENSE))

## Requisitos
- NPM 6.9.0 , NodeJs 12.0

## Guía de instalación
1. Clonar el repositorio
```shell
git clone https://github.com/cesarsasi/visor-mapas-frontend
```

2. Ingresar a la carpeta del repositorio e instalar dependencias
```shell
cd visor-mapas-frontend
npm i
```

3. Configurar URL de las variables dentro de los archivos ocultos .env.development y .env.production
```shell
# URL base de API rest del backend
VUE_APP_API_URL=http://localhost:2000/ 
# URL base de GeoServer
VUE_APP_GEOSERVER_URL=http://localhost:8600/geoserver/
# URL base de Nominatim
VUE_APP_API_NOMINATIM=http://localhost:7070/
# URL de maptiles propios (no OpenStreetMap)
VUE_APP_MAPTILES=http://localhost:82/{z}/{x}/{y}.jpg
# Nombre del espacio de trabajo por defecto utilizado para subir capas a GeoServer (debe estar configurado en GeoServer)
VUE_APP_DEFAULT_WORKSPACE=visor_upload
```

4. Guardar los cambios en los archivos y ejecutar en visor de mapas
- Ejecutar directamente (development)
```shell
npm run serve
```

- Compilar (production)
```shell
npm run build
# Archivos de salida se encuentran en ./dist
# Pueden ser alojados en un servidor Nginx
```

## Estructura del código

#### src/views/
Vistas del visor. Se didiven en los siguientes grupos:
- src/views/account/: Vistas relacionadas con la administración de la cuenta del usuario
- src/views/categories/: Gestión de categorías de las capas
- src/views/geoprocessings/: Administración de geoprocesos
- src/views/groups/: Gestión de los grupos de usuarios del visor
- src/views/layers/: Administración de las capas
- src/views/maps/: Administración de los mapas
- src/views/users/: Administración de los usuarios
- src/views/visor/: Vistas donde se muestra el mapa

#### src/components/
Componentes del mapa
- src/components/mapControls/: Controles del mapa (zoom,nav drawer, simbología del mapa)
- src/components/mapTools/: Herramientas del mapa. Se dividen en carpetas de acuerdo a cada herramienta. Cada herramienta está compuesta por al menos dos archivos: 
    - Archivo Vue que corresponde a la interfaz gráfica con la cual el usuario interactúa y se muestra sobre el mapa
    - Archivo JS que es un mixin, el cual se importa directamente en el mapa para poder realizar los cambios en el mapa.
    Las herramientas son las siguientes:
    1. baseMap/: Herramienta para el cambio del mapa base.
    2. charts/: Herramienta que despliega una capa con features que pueden ser seleccionadas por el usuario y mostrar información estadística mediante gráficos de acuerdo a la feature seleccionada.
    3. export/: Herramienta de exportación del mapa, mediante screenshot o PDF.
    4. geoprocessings/: Herramienta de ejecución de geoprocesos mediante la API de Héctor.
    5. help/: Herramienta de ayuda del mapa al usuario.
    6. layers/: Lista de capas que se muestran en el mapa.
    7. measures/: Herramienta de medición del mapa.
    8. search/: Buscador de ubicaciones mediante una dirección en el mapa.
    9. select/: Herramienta de selección de punto y área en el mapa.
- src/components/mixin/: mixins utilizados en la aplicación. Contiene código que es utilizado por más de un componente. Aquí se encuentran las consultas rest utilizadas en las vistas y los métodos para obtener y cambiar el estado de las herramientas del mapa (vuex.js).

#### src/store.js
Se utiliza Vuex para el manejo del estado y comunicación de las herramientas del mapa.

#### Como cambiar el estado de una herramienta
Es necesario que el componente o mixin incluya el mixin vuex.js
1. Seleccionar una herramienta
```js
this.selectMapTool('nombre_herramienta', true)
```
2. Desactivar una herramienta
```js
this.selectMapTool('nombre_herramienta', false)
```
3. Realizar acción en una herramienta
```js
this.executeMapToolAction('nombre_herramienta', 'accion_a_realizar', {parametro: 'valor'})
```
4. Escuchar acción de una herramienta en otro compoente
```js
watch: {
    'mapToolExecuted': function (newValue, oldValue) {
        // newValue coresponde al nombre de la herramienta  que realizó una acción
        switch (newValue) {
            case 'nombre_herramienta':
                if(this.mapToolValues('nombre_herramienta').action == 'accion_a_realizar') {
                    if(this.mapToolValues('nombre_herramienta').data.parametro == 'valor') {
                        // Acción a realizar
                    }
                }
                break;
            default:
                break;
        }
    }
}
```
Si se agregan nuevas herramientas, debe agregarse a store.js y realizar los cambios necesarios para que el estado y la comunicación de la herramienta se haga correctamente.

#### Componentes importantes
- src/components/MapContainer.vue: Es la vista completa del mapa. Incluye al mapa y las herramientas que se muestran sobre él.
- src/components/MapView.vue: Es el mapa como tal (el mismo que se incluye en MapContainer). Este componente escucha las acciones ejecutadas por las herramientas y realiza las respectivas acciones y cambios en el mapa.
- src/components/Login.vue: Este componente se usa para el inicio de sesión. Guarda la cookie de permisos del usuario, utilizadas para permitir el acceso a las distintas vistas del visor de mapas de acuerdo a los permisos (la cookie de sesión del usuario se almacena automáticamente). En caso que la cookie de sesión expire, el comportamiento está definido en main.js. El cierre de sesión está en src/views/Logout.vue.
- .env.development y .env.production: variables de entorno para desarrollo y producción, respectivamente. Contiene las siguientes variables:
    - VUE_APP_API_URL: URL donde se encuentra corriendo la API backend.
    - VUE_APP_GEOSERVER_URL: URL de la instancia de GeoServer en la cual se encuentran las capas.
    - VUE_APP_API_NOMINATIM: URL de la API de geocodificación de Nominatim
    - VUE_APP_MAPTILES: URL donde se encuentra el servidor de capas alternativas para el cambio de mapa base. Debe estar en formato XYZ.
    - VUE_APP_DEFAULT_WORKSPACE: Espacio de trabajo de GeoServer donde las capas se suben desde el visor de mapas. Para un correcto funcionamiento, se recomienda que tanto el espacio de trabajo, almacen de datos y schema de la base de datos en GeoServer tengan el mismo nombre.
A continuación se muestra un ejemplo de los archivos de configuración. Notar que todas las URL terminan con un '/'.

```bash
VUE_APP_API_URL=http://localhost:2000/
VUE_APP_GEOSERVER_URL=http://localhost:8600/geoserver/
VUE_APP_API_NOMINATIM=http://localhost:7070/
VUE_APP_MAPTILES=http://localhost:82/{z}/{x}/{y}.jpg
VUE_APP_DEFAULT_WORKSPACE=visor_upload
``` 

## Licencia
2-Clause BSD License
