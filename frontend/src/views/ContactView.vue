<template>
  <div class="about">
    <h1>Kontaktný formulár</h1>

    <form action="https://api.web3forms.com/submit" method="POST">
      <input type="hidden" name="access_key" value="37a1ac8d-4ab9-4c34-914e-38aa237c7f14">

      <input type="text" placeholder="Vaše meno" name="name" required>

      <input type="email" placeholder="Váš email" name="email" required>
      <textarea name="message"  placeholder="Popis"  required></textarea>
      <br>
      <button @click="postData">Poslať</button>

    </form>
  </div>

  <div class="contact-info">
    <br>
    <h2>Kontakt:</h2>
    <p><a class="link" href="tel:+421911421521">+421 911 421 521</a></p>

    <a class="link" :href="'mailto:' + 'jozko.mrkvicka@email.com'">jozko.mrkvicka@email.com</a>
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

<style scoped>
textarea  {
  padding: 8px;
  height: 225px;
  margin-top: 25px;
  width: 70%;
  border-radius: 15px;
}

input {
  margin-top: 25px;
  padding: 8px;

  width: 70%;
  border-radius: 15px;
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
  background-color: #f2cc8f;
  color: #5f0f40;
  border: none;
  cursor: pointer;
  margin-top: 10px;
  border-radius: 15px;
  font-size: 17px;
  height: 40px;
  width: 100px;
  font-weight: 1000;
}

button + p {
  margin-top: 16px;
}

.about {
  position: fixed;
  top: 180px;
  left: 20px;
  width: 50%;

  z-index: 1000;
  padding: 10px;
}

.contact-info {
  right: 150px;
  position: fixed;
  top: 180px;

}

button:hover {
  background-color: #ddb892;
}

h1 {
  font-weight: 700;
}

h2 {
  font-weight: 700;
}

a {
  font-weight: 500;
}

p {
  font-weight: 500;
}

iframe{
  height: 400px;
}

.link{
  color: #F7CFD8;
}
</style>
