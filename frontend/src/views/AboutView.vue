<script setup>
import axios from 'axios';
import {ref} from 'vue';

const output = ref('');
const name = ref('');
const description = ref('');
const score = ref(0);
const errorMsg = ref(0);
const successMsg = ref(0);
const isSubmitted = ref(false);
const reviews = ref([])

async function fetchData() {
  try {
    const response = await axios.get('http://localhost:8000/review');
    console.log(response.data);
    console.log('Fetched reviews:', response.data);
    reviews.value = response.data;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}

async function postData() {
  successMsg.value = '';
  errorMsg.value = '';
  isSubmitted.value = true;
  if (!name.value || score.value === null || score.value < 1 || score.value >5) {
    return;
  }

  try {
    const response = await axios.post('http://localhost:8000/review', {
      name: name.value,
      description: description.value,
      score: score.value
    });

    output.value = response.data.message;
    successMsg.value = 'Recenzia bola úspešne pridaná.';
    errorMsg.value = '';
    fetchData();
  } catch (error) {
    successMsg.value = '';
    console.error('Error posting data:', error);
    errorMsg.value = error.response.data.message;
  }
}
fetchData();
</script>



<template>
  <div class="about">
    <h1>Pridať recenziu</h1>

    <div class="input-window">
      <label for="name">Meno:</label>
      <br>
      <input id="name" v-model="name" type="text"/>
      <br>
      <span v-if="isSubmitted && !name" style="color: red;">Zadajte meno.</span>
    </div>

    <div class="input-window">
      <label for="score">Hodnotenie: (1-5)</label>
      <br>
      <input id="score" v-model="score" type="number"/>
      <br>
      <span v-if="(score < 1 || score> 5 || !score) && isSubmitted" style="color: red;">Zadajte hodnotenie v rozmedzí 1-5.</span>
    </div>

    <div class="input-window">
      <label for="description">Popis:</label>
      <br>
      <input id="description" v-model="description" type="text" />
    </div>

    <button @click="postData">Poslať</button>


    <p v-if="errorMsg" style="color: red;">{{ errorMsg }}</p>
    <p v-if="successMsg" style="color: green;">{{ successMsg }}</p>

    <br>
    <br>

  </div>
  <div class="show-reviews">
    <h2>Existing Reviews</h2>
    <div v-if="reviews.length > 0">
      <ul>
        <li v-for="review in reviews" :key="review.reviewID">
          <strong>{{ review.name }} {{ review.score }}:</strong> {{ review.description }}
        </li>
      </ul>
    </div>
    <div v-else>
      <p>No reviews yet.</p>
    </div>
  </div>

</template>

<style scoped>

.input-window {
  margin-bottom: 16px;
}

.show-reviews{
  top: 180px;
}

input {
  padding: 8px;
  margin-top: 4px;
  width: 70%;
  border-radius: 8px;
}

span {
  font-size: 15px;
  margin-left: 5px;
}

.filled input:invalid {
  border: 2px solid red;
  outline: 2px solid red;
}


button {
  padding: 10px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
  margin-top: 16px;
  border-radius: 8px;
}


button + p {
  margin-top: 16px;
}

.about {
  position: fixed;
  top: 180px;
  left: 0;
  width: 40%;

  z-index: 1000;
  padding: 10px;
}
</style>
