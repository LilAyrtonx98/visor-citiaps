<template>
  <div class="statContainer" :class="{statContainerAlt : chartContainerAlt}">
    <p class="stat-title">{{title}}</p>
    <p class="stat-subtitle">{{subtitle}}</p>
    <div v-for="(stat, i) in stats" v-bind:key="i">
      <span class="stat-number">{{' ' + stat.number}}</span>
      <span class="stat-label">{{' ' + stat.label}}</span>
    </div>
  </div>
</template>

<style lang="scss" scoped>

</style>


<script>
export default{
  props:['stats', 'title', 'subtitle', 'chartContainerAlt'],
  data(){
    return {
        
    }
  },
  computed: {

  },
  methods: {

  },
  created(){
  
  },
  watch: {

  }
}
</script>