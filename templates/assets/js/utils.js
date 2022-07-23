// form.serialize() to json
function formSerializeToJson(data) {
  data = decodeURIComponent(data); // decode URI to human string
  data = data.replace(/&/g, '","');
  data = data.replace(/=/g, '":"');
  data = '{"' + data + '"}';
  return data;
}

function FormToJson(form) {
  let data = $(form).serialize();
  let json_data = formSerializeToJson(data);
  //   console.log(data);
  //   console.log(json_data);
  return json_data;
}
