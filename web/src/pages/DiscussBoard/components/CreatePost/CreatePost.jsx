import React from 'react';
import { Form, Input, Select } from 'antd';
import * as $u from '../../services/DiscussService';
import { connect } from 'react-redux'
import { ACTION_TYPES } from '../../../../redux/actions/actionType';
import { topicActions } from '../../../../redux/actions/createTopic';

const layout = {
    labelCol: { span: 4 },
    wrapperCol: { span: 16 },
};

const { Option } = Select;

const validateMessages = {
    required: '${label} is required!'    
};

const PAYLOAD = {
    title: 'postName',
    topic: 'subredditName',
    url: 'url',
    description: 'description'
};

class CreatePost extends React.Component {    
    componentDidMount() {                
        this.props.getAllTopics();
    }

    /** e is the input element ref or id of post (1-based) */
    onChange = (e, field) => {
        let val = '';
        let isValid = true;
        if (field === 'topic') { // dropdown field type
            const idx = this.props.topic.topicList.length - e; // reversed idx
            val = this.props.topic.topicList[idx].name;            
        } else { // input field type
            val = e.target.value;            
        }

        if (this.props.post[field].isRequired) {
            isValid = val != null && val != undefined && val != '';
        }

        const data = {
            ...this.props.post,
            payload: {
                ...this.props.post.payload,
                [PAYLOAD[field]]: val
            },
            [field]: {...this.props.post[field], isValid, value: val}
        }
        this.props.dispatch({ type: ACTION_TYPES.CREATE_POST,  data });   
    };

    render() {
        return (
            <Form {...layout} name="nest-messages" validateMessages={validateMessages}>
                <Form.Item name='title' label="Title" rules={[{ required: true }]}>
                    <Input onChange={(e) => this.onChange(e, 'title')}/>
                </Form.Item>
                <Form.Item name='url' label="URL">
                    <Input onChange={(e) => this.onChange(e, 'url')} />
                </Form.Item>                        
                <Form.Item name='topic' label="Select Topic" rules={[{ required: true }]}>
                    <Select
                        showSearch
                        placeholder="Select a Topic"
                        optionFilterProp="name"
                        onChange={(e) => this.onChange(e, 'topic')}
                        options={this.props.topic.topicList.map((topic) => { return { label: topic.name, value: topic.id } })}                                         
                        filterOption={(input, option) =>
                            option.label.toLowerCase().indexOf(input.toLowerCase()) >= 0}>                            
                    </Select>
                </Form.Item>
                <Form.Item name='description' label="Description">
                    <Input.TextArea onChange={(e) => this.onChange(e, 'description')} />
                </Form.Item>                
            </Form>            
        );
    }
}

const mapStateToProps = (state) => {
    return {
        post: state.postReducer,
        topic: state.getAllTopicReducer
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        getAllTopics: () => dispatch(topicActions.getTopicList()),
        dispatch
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(CreatePost);