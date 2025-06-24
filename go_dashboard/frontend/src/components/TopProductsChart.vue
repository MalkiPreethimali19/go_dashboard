<template>
  <div class="card">
    <h2>Top 20 Frequently Purchased Products</h2>
    <div class="chart-wrapper">
      <Bar :data="chartData" :options="chartOptions" v-if="loaded" />
    </div>
  </div>
</template>

<script>
import { Bar } from 'vue-chartjs'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
} from 'chart.js'
import api from '../api'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

export default {
  components: { Bar },
  data() {
    return {
      chartData: {
        labels: [],
        datasets: []
      },
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
            tooltip: {
                callbacks: {
                    label: (context) => {
                        const qty = context.raw
                        const stock = context.chart.data.datasets[0].stockQuantities[context.dataIndex]
                        return `Quantity Sold: ${qty} | Stock Left: ${stock}`
                    }
                }
            },
            legend: { display: false }
        },
        scales: {
          x: { 
            autoSkip: false,
            font: { size: 12 }
           },
          y: { beginAtZero: true }
        }
      },
      loaded: false
    }
  },
  mounted() {
    api.get('/top-products').then((res) => {
      const labels = res.data.map(p => p.product_name)
      const values = res.data.map(p => p.total_quantity)
      const stockQuantities = res.data.map(p => p.stock_quantity)

      this.chartData = {
        labels,
        datasets: [{
          label: 'Quantity Sold',
          data: values,
          backgroundColor: this.generateColors(values.length),
          stockQuantities
        }]
      }
      this.loaded = true
    }).catch(err => {
      console.error("Error loading top products:", err)
    })
  },
  methods: {
    generateColors(count) {
      const colors = []
      for (let i = 0; i < count; i++) {
        colors.push(`hsl(${(i * 360) / count}, 70%, 60%)`)
      }
      return colors
    }
  }
  
}
</script>

<style scoped>
.card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
}

.chart-wrapper {
  width: 100%;
  overflow-x: auto;
  height: 500px; 
}
</style>

