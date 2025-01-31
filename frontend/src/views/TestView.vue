<script setup>
import axios from 'axios';
import { ref } from 'vue';

const output = ref('')

let bad = 0;
const good = ref(0);

const inputData = ref('')

async function fetchData() {
  const response = await axios.get('http://localhost:8000/');
  console.log(response.data);
  output.value = response.data;
}

async function postData() {
  if (inputData.value.length === 0) {

    return
  }
  const response = await axios.post('http://localhost:8000/write', {
    data: inputData.value
  });
  console.log(response.data);
  output.value = response.data;
}

fetchData();
</script>

<template>
  <div>
    <br>
    <h2>Kontakt:</h2>
    <p><a href="tel:+421911421521">+421 911 421 521</a></p>

    <a :href="'mailto:' + 'jozko.mrkvicka@email.com'">jozko.mrkvicka@email.com</a>
    <br>

    <br>
    <h2>Adresa:</h2>
    <p>Mašličková 13</p>
    <p>Brno - Nový Lískovec</p>
    <p>634 00</p>
    <br>

    <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d5216.374822106611!2d16.566499058423364!3d49.178030265998835!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x471295bf39ca282b%3A0xefba8968c501c195!2sFakultn%C3%A1%20nemocnica%20Brno!5e0!3m2!1ssk!2scz!4v1737492722089!5m2!1ssk!2scz" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>


  </div>
</template>
