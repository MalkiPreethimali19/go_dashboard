<template>
  <div class="card">
    <h2 class="title">Country-Level Revenue</h2>

    <div class="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>Country</th>
            <th>Product</th>
            <th>Revenue</th>
            <th>Transactions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in data" :key="row.product_name + row.country">
            <td>{{ row.country }}</td>
            <td>{{ row.product_name }}</td>
            <td>${{ row.total_revenue.toFixed(2) }}</td>
            <td>{{ row.transaction_count }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import api from '../api';

export default {
  data() {
    return {
      data: []
    };
  },
  mounted() {
    api.get('/country-revenue')
      .then(res => {
        this.data = res.data;
      })
      .catch(err => {
        console.error("API error:", err);
      });
  }
};
</script>

<style scoped>
.card {
  background-color: #fff;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.05);
  margin-bottom: 2rem;
}

.title {
  font-size: 1.4rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: #333;
}

/* Scrollable wrapper for both directions */
.table-wrapper {
  max-height: 400px;
  overflow-x: auto;
  overflow-y: auto;
  border: 1px solid #ddd;
  border-radius: 6px;
}

/* Table styles */
table {
  width: 100%;
  border-collapse: collapse;
  min-width: 600px;
}

th, td {
  padding: 12px 16px;
  text-align: left;
  border-bottom: 1px solid #eee;
  font-size: 0.95rem;
}

th {
  background-color: #f4f4f4;
  font-weight: 600;
}

tr:hover {
  background-color: #f9f9f9;
}

/* Responsive tweaks */
@media (max-width: 600px) {
  th, td {
    padding: 10px;
    font-size: 0.85rem;
  }

  .title {
    font-size: 1.2rem;
  }
}
</style>
