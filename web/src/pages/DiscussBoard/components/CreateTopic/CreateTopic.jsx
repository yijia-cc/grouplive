import React from 'react';
import { Form, Input } from 'antd';
import { connect } from 'react-redux'
import { ACTION_TYPES } from '../../../../redux/actions/actionType';
import './CreateTopic.css';

const layout = {
    labelCol: { span: 6 },
    wrapperCol: { span: 14 },
};

const validateMessages = {
    required: '${label} is required!'
};

const PAYLOAD = {
    title: 'name',
    description: 'description'
};

class CreateTopic extends React.Component {
    onChange = (val, field) => {
        let isValid = true;
        if (this.props.topic[field].isRequired) {
            isValid = val != null && val != undefined && val != '';
        }

        const data = {
            ...this.props.topic,
            payload: {
                ...this.props.topic.payload,
                [PAYLOAD[field]]: val
            },
            [field]: {...this.props.topic[field], isValid, value: val}
        }
        this.props.dispatch({ type: ACTION_TYPES.CREATE_TOPIC,  data});   
    };

    render() {
        return (
            <Form {...layout} name="nest-messages" validateMessages={validateMessages}>
                <Form.Item name='Title' label="Title" rules={[{ required: true }]}>
                    <Input onChange={(e) => this.onChange(e.target.value, 'title')}/>
                </Form.Item>                       
                <Form.Item name='description' label="Description">
                    <Input.TextArea onChange={(e) => this.onChange(e.target.value, 'description')}/>
                </Form.Item>
            </Form>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        topic: state.topicReducer
    }
}

export default connect(mapStateToProps)(CreateTopic);;