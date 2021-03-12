<template>
  <div class="container">
    <div class="inner">
      <h1>Services discovered</h1>

      <div style="display: flex;">
        <div style="display: flex; flex-direction: column; align-items: center;">
          <span style="margin-bottom: 8px;">Service/Ingress per Capability</span>
          <div style="position: relative; width:600px;">
            <CapabilitiesServicesChart v-bind:chartdata="this.generateChartDataForServicesPerCapability" :options="{responsive: true, maintainAspectRatio: true, legend: {
              display: false
            },
            tooltips: {
              titleFontSize: 18,
              bodyFontSize: 18,
            }
            }" />
          </div>
        </div>

        <div style="display: flex; flex-direction: column; align-items: center; margin-left: 15px;">
          <span style="margin-bottom: 8px;">Services vs Ingress</span>
          <div style="position: relative; width:600px;">
            <CapabilitiesServicesChart v-bind:chartdata="this.generateChartDataForServiceVsIngress" :options="{responsive: true, maintainAspectRatio: true, legend: {
              display: false
            },
            tooltips: {
              titleFontSize: 18,
              bodyFontSize: 18,
            }
            
            }" />
          </div>
        </div>        
      </div>

      
      <CapabilityListing v-bind:capabilities="this.capabilities" v-bind:services="this.services" v-bind:costs="this.costs" v-bind:logged_in="this.logged_in" />
    </div>
  </div>
</template>

<script>
import CapabilityListing from './CapabilityListing.vue'
import CapabilitiesServicesChart from './CapabilitiesServicesChart.js';


export default {
  name: 'HelloWorld',
  props: {
    msg: String,
    logged_in: Boolean,
    capabilities: Array,
    costs: {},
    services: {}
  },
  data() {
    return {
    }
  },
  components: {
    CapabilityListing,
    CapabilitiesServicesChart
  },

  computed: {
    generateChartDataForServicesPerCapability() {
        var data = [];
        var colours = [];
        var labels = [];
        var services = this.services;
        Object.keys(services).forEach(namespace => {
          data.push(Object.keys(services[namespace]).length);
          colours.push(this.dynamicColour());
          labels.push(this.getCapability(namespace));
        });

        var payload = {
          datasets: [{
            data: data,
            backgroundColor: colours
          }],
          labels: labels
        };

        return payload;      
    },

    generateChartDataForServiceVsIngress() {
      var ingressCount = 0;
      var serviceCount = 0;
      var services = this.services;

      Object.values(services).forEach(namespace => {
        Object.values(namespace).forEach(obj => {
          if (obj.Kind.valueOf() === "Ingress") {
            ingressCount += 1;
          } else {
            serviceCount += 1;
          }
        });
      });

      var payload = {
        datasets: [{
          data: [ingressCount, serviceCount],
          backgroundColor: [this.dynamicColour(), this.dynamicColour()]
        }],
        labels: ["Ingress", "Service"]
      };

        return payload;      
    }
  },

  methods: {
    dynamicColour() {
      var r = Math.floor(Math.random() * 255);
      var g = Math.floor(Math.random() * 255);
      var b = Math.floor(Math.random() * 255);
      return "rgb(" + r + "," + g + "," + b + ")";
    },

    getCapability(rootid) {
      var result = this.capabilities.filter(cap => {
        if (cap.rootId.valueOf() === rootid.valueOf()) {
          return true;
        }
        return false;
      })

      if (result.length > 0) {
        return result[0].name;
      }

      return rootid;
    }
  }
}


</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

h1 {
  font-size: 4em;
  margin-left: 15px;
}

.container {
    display: flex;
    flex-direction: column;
    margin: auto;
    width: 70%;
    align-items: flex-start;
    background-color: #423e48;
    color: #c9c9c9;
    padding-bottom: 15px;
}

.inner {
    padding-top: 35px;
    display: flex;
    flex-direction: column;
    width: 100%;
}
</style>
