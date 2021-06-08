import React from "react";
import { Form, Input, Button, message } from "antd";
import axios from "axios";
import { BASE_URL } from "../../constants";
import "./index.css";

const formItemLayout = {
  labelCol: {
    xs: { span: 24 },
    sm: { span: 8 },
  },
  wrapperCol: {
    xs: { span: 24 },
    sm: { span: 16 },
  },
};
const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 16,
      offset: 0,
    },
    sm: {
      span: 16,
      offset: 8,
    },
  },
};

function Register(props) {
  const [form] = Form.useForm();

  const onFinish = (values) => {
    const {
      username,
      password,
      firstname,
      lastname,
      email,
      address,
      apt_number,
      phone,
    } = values;
    const opt = {
      method: "POST",
      url: `${BASE_URL}/sign-up`,
      data: {
        username,
        password,
        firstname,
        lastname,
        email,
        address,
        apt_number,
        phone,
      },
      headers: { "content-type": "application/json" },
    };

    axios(opt)
      .then((response) => {
        console.log(response);
        // case1: registered success
        if (response.status === 200) {
          message.success("Registration succeed!");
          props.history.push("/login");
        }
      })
      .catch((error) => {
        console.log("register failed: ", error.message);
        message.success("Registration failed!");
        // throw new Error('Signup Failed!')
      });
  };

  return (
    <Form
      {...formItemLayout}
      form={form}
      name="register"
      onFinish={onFinish}
      className="register"
    >
      <Form.Item
        name="username"
        label="Username"
        rules={[
          {
            required: true,
            message: "Please input your Username!",
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="email"
        label="Email"
        rules={[
          {
            required: true,
            message: "Please input your Username!",
          },
          { type: "email" },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="firstname"
        label="FirsetName"
        rules={[
          {
            required: true,
            message: "Please input your firsetName!",
            whitespace: true,
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="lastname"
        label="LastName"
        rules={[
          {
            required: true,
            message: "Please input your lastName!",
            whitespace: true,
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="phone"
        label="Phone Number"
        rules={[{ required: true, message: "Please input your phone number!" }]}
      >
        <Input style={{ width: "100%" }} />
      </Form.Item>
      <Form.Item
        name="apt_number"
        label="Apartment Number"
        rules={[
          {
            required: true,
            message: "Please input your Apartment Number!",
            whitespace: true,
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="address"
        label="Address"
        rules={[
          {
            required: true,
            message: "Please input your address!",
            whitespace: true,
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="password"
        label="Password"
        tooltip="Password length must be at least 8 characters"
        rules={[
          {
            required: true,
            message: "Please input your password!",
          },
          {
            min: 8,
            message: "Password length must be at least 8 characters",
          },
        ]}
        hasFeedback
      >
        <Input.Password />
      </Form.Item>

      <Form.Item
        name="confirm"
        label="Confirm Password"
        dependencies={["password"]}
        hasFeedback
        rules={[
          {
            required: true,
            message: "Please confirm your password!",
          },
          ({ getFieldValue }) => ({
            validator(rule, value) {
              if (!value || getFieldValue("password") === value) {
                return Promise.resolve();
              }
              return Promise.reject(
                "The two passwords that you entered do not match!"
              );
            },
          }),
        ]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item {...tailFormItemLayout}>
        <Button type="primary" htmlType="submit" className="register-btn">
          Register
        </Button>
      </Form.Item>
    </Form>
  );
}

export default Register;
