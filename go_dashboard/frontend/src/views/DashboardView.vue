<template>
  <div class="dashboard">
    <CountryRevenueTable />
    <TopProductsChart />
    <MonthlySalesChart />
    <TopRegionsChart />
  </div>
</template>

<script>
import CountryRevenueTable from '@/components/CountryRevenueTable.vue'
import TopProductsChart from '@/components/TopProductsChart.vue'
import MonthlySalesChart from '@/components/MonthlySalesChart.vue'
import TopRegionsChart from '@/components/TopRegionsChart.vue'

export default {
  components: {
    CountryRevenueTable,
    TopProductsChart,
    MonthlySalesChart,
    TopRegionsChart
  }
}

</script>

<style scoped>
.dashboard {
  display: grid;
  grid-template-columns: 1fr; /* 1 column by default */
  gap: 1.5rem;
  padding: 1.5rem;
}

/* On medium+ screens: 2 columns */
@media (min-width: 768px) {
  .dashboard {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
