<script setup>
import axios from 'axios';
import {ref} from 'vue';

const output = ref('');
const name = ref('');
const description = ref('');
const score = ref(0);
const successMsg = ref(0);
const isSubmitted = ref(false);
const reviews = ref([])

const avg_score = ref(0.0)


async function fetchData() {
  try {
    const response = await axios.get('http://localhost:8000/review');
    console.log(response.data);
    console.log('Fetched reviews:', response.data);
    reviews.value = response.data;
    calculateAvgScore();
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}

function calculateAvgScore() {
  if (reviews.value.length > 0) {
    let totalScore = 0;
    for (let i = 0; i < reviews.value.length; i++) {
      totalScore += reviews.value[i].score;
    }
    avg_score.value = (totalScore / reviews.value.length).toFixed(1);
  } else {
    avg_score.value = 0;
  }
}

async function postData() {
  successMsg.value = '';
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
    fetchData();
  } catch (error) {
    successMsg.value = '';
    console.error('Error posting data:', error);
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
      <input id="name" v-model="name" type="text" :class="{ invalid: isSubmitted && !name }"/>
      <br>
      <span v-if="isSubmitted && !name" style="color: #E195AB;">Zadajte meno.</span>
    </div>

    <div class="input-window">
      <label for="score">Hodnotenie: (1-5)</label>
      <br>
      <input id="score" v-model="score" type="number" :class="{ invalid: isSubmitted && (score < 1 || score > 5 || !score) }"/>
      <br>
      <span v-if="(score < 1 || score> 5 || !score) && isSubmitted" style="color: #E195AB;">Zadajte hodnotenie v rozmedzí 1-5.</span>
    </div>

    <div class="input-window">
      <label for="description">Popis:</label>
      <br>
      <textarea class="description-input" id="description" v-model="description" type="text"></textarea>
    </div>

    <button @click="postData">Poslať</button>
    <p v-if="successMsg" style="color: #80ed99;">{{ successMsg }}</p>
  </div>


  <div class="show-reviews">
    <h1>Recenzie: {{ avg_score }}</h1>
    <hr class="first-line">

    <div v-if="reviews.length > 0">
      <ul class="review-list">
        <li v-for="review in reviews" :key="review.reviewID">
          <strong>
            {{ review.name }}: {{ review.score }}
          </strong>
          <span v-if="review.description"></span>
          <span class="descript" v-if="review.description">{{ review.description }}</span>
          <hr class="custom-line">
        </li>
      </ul>
    </div>
    <div v-else>
      <p>No reviews yet.</p>
    </div>
  </div>

</template>

<style scoped>

.first-line {
  border: none;
  height: 3px;
  background-color: #5f0f40;
  width: 100%;
  margin: 10px auto;
}

.custom-line {
  border: none;
  height: 2px;
  background-color: #5f0f40;
  width: 100%;
  margin: 10px auto;
}

h1 {
  font-weight: 700;
}
strong  {
  font-size: 25px;
  color: #240046;
  background-color: #e6ccb2;
  font-weight: 700;
  padding: 5px;
  border-radius: 10px;
}

.input-window {
  margin-bottom: 16px;
}

.review-list {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.review-list li {
  margin-left: 10px;
  margin-bottom: 10px;
}

.show-reviews{
  position: fixed;
  color: #5f0f40;
  border-radius: 20px;
  padding: 15px;
  width: 50%;
  height: 620px;
  background-color: white;
  top: 200px;
  right: 120px;
  overflow-y: auto;
}

input {
  padding: 8px;
  margin-top: 4px;
  width: 80%;
  border-radius: 15px;
}

textarea {
  padding: 8px;
  margin-top: 4px;
  height: 200px;
  width: 80%;
  border-radius: 15px;
}

span {
  font-size: 20px;
  margin-left: 15px;
  display: block;
  font-weight: 600;
}

.invalid {
  border: 1px solid #E195AB;
  outline: 1px solid #E195AB;
}

button {
  padding: 10px 20px;
  background-color: #f2cc8f;
  color: #5f0f40;
  font-weight: 1000;
  border: none;
  cursor: pointer;
  margin-top: 10px;
  border-radius: 15px;
  font-size: 17px;
  height: 40px;
  width: 100px;
}
button:hover {
  background-color: #ddb892;
}

button + p {
  margin-top: 16px;
}

.about {
  position: fixed;
  top: 180px;
  left: 20px;
  width: 40%;

  z-index: 1000;
  padding: 10px;
}

label {
  font-weight: 500;
}
</style>
