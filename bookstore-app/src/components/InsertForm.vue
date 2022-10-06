<template>
  <div class="form-bg">
    <div class="container">
      <div class="row">
        <div class="col-md-offset-3 col-md-6 col-sm-offset-3 col-sm-6 myform">
          <div class="form-container">
            <div class="form-icon">
              <span><i class="fa fa-file-alt"></i></span>
            </div>
            <form class="form-horizontal">
              <h3 class="title">Insert Book details</h3>
              <div class="form-group">
                <input
                  v-model="title"
                  type="text"
                  class="form-control"
                  placeholder="Title"
                  required
                />
              </div>
              <div class="form-group">
                <input
                  v-model="author"
                  type="text"
                  class="form-control"
                  placeholder="Author"
                  required
                />
              </div>
              <div class="form-group">
                <input
                  v-model="isbn"
                  type="text"
                  class="form-control"
                  placeholder="ISBN"
                  required
                />
              </div>
              <div class="form-group">
                <input
                  v-model="edition"
                  type="text"
                  class="form-control"
                  placeholder="Edition"
                  required
                />
              </div>
              <div class="form-group">
                <input
                  v-model="year"
                  type="number"
                  class="form-control"
                  placeholder="Year"
                  required
                />
              </div>
              <div class="form-group">
                <input
                  v-model="image"
                  type="link"
                  class="form-control"
                  placeholder="Image Link"
                  required
                />
              </div>
              <div class="form-group">
                <button class="btn insert">Add Book</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Button",
  data: function () {
    return {
      title: "",
      author: "",
      isbn: "",
      edition: "",
      year: "",
      image: "",
    };
  },
  methods:{
      submitForm(){
        title = this.title
      }
  },
};
</script>

<style>
.form-container {
  font-family: "Open Sans", sans-serif;
  text-align: center;
}
.form-container .form-icon {
  color: #fff;
  background: linear-gradient(to right, #0f3710, #feb101);
  font-size: 50px;
  line-height: 110px;
  width: 110px;
  height: 110px;
  margin: 0 auto -55px;
  border-radius: 50%;
  position: relative;
  z-index: 1;
}
.form-container .form-horizontal {
  background: linear-gradient(45deg, #fff 50%, #f9f9f9 50%);
  background-size: 20px 20px;
  padding: 75px 40px 15px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
}
.form-horizontal .title {
  color: #4f4f4f;
  font-size: 20px;
  text-transform: uppercase;
  padding: 0 25px;
  margin: 0 0 20px;
  display: inline-block;
  position: relative;
}
.form-horizontal .title:before,
.form-horizontal .title:after {
  content: "";
  background: #4f4f4f;
  width: 15px;
  height: 2px;
  transform: translateY(-50%);
  position: absolute;
  top: 50%;
  left: 0;
}
.form-horizontal .title:after {
  left: auto;
  right: 0;
}
.form-horizontal .form-group {
  margin: 0 0 15px;
}
.form-horizontal .form-group:last-of-type {
  border-top: 2px solid #e5e5e5;
  padding: 15px 15px 0;
  margin: 0 -40px 0;
}
.form-horizontal .form-control {
  background: #e6e6e6;
  font-size: 16px;
  text-align: center;
  letter-spacing: 1px;
  height: 40px;
  padding: 6px 30px;
  border-radius: 8px;
  border: none;
  box-shadow: none;
  border-top: 4px solid #cecece;
  border-left: 4px solid #cecece;
}
.form-horizontal .form-control:focus {
  box-shadow: none;
}
.form-horizontal .form-control::placeholder {
  color: #6c6c6c;
}
.form-horizontal .insert {
  color: white;
  background-color: #cfa524;
  font-size: 20px;
  font-weight: 600;
  text-transform: uppercase;
  width: 100%;
  padding: 6px 10px;
  border: none;
  border-radius: 8px;
  box-shadow: 0 -3px 0 rgba(0, 0, 0, 0.1) inset;
  display: block;
  transition: all 0.3s ease 0s;
}
.form-horizontal .insert:hover,
.form-horizontal .insert:focus {
  color: white;
  box-shadow: 0 3px 0 rgba(0, 0, 0, 0.1) inset, 0 0 10px rgba(0, 0, 0, 0.1);
  background: linear-gradient(to right, #0f3710, #feb101);
}
@media only screen and (min-width: 800px) {
  .myform {
    width: 500px;
    margin-left: 25%;
    padding-bottom: 5%;
  }
}
@media only screen and (max-width: 359px) {
  .form-container .form-horizontal {
    padding-left: 20px;
    padding-right: 20px;
  }
  .form-horizontal .form-group:last-of-type {
    padding: 15px 15px 0;
    margin: 0 -20px 0;
  }
}
</style>