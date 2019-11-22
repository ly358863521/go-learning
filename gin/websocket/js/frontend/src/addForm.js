import React from 'react';

import { Form, Icon, Input, Button } from 'antd';

function hasErrors(fieldsError) {
  return Object.keys(fieldsError).some(field => fieldsError[field]);
}

class AddForm extends React.Component {
  componentDidMount() {
    // To disabled submit button at the beginning.
    this.props.form.validateFields();
  }

  handleSubmit = e => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log('Received values of form: ', values);
        this.outsideSubmit(values)
      }
    });
    return false;
  };

  outsideSubmit = ()=>{console.log("undef submit;")}

  render() {
    const { getFieldDecorator, getFieldsError, getFieldError, isFieldTouched } = this.props.form;

    const { handleSubmit } = this.props;

    this.outsideSubmit = handleSubmit;
    // Only show error after a field is touched.
    const NameError = isFieldTouched('Name') && getFieldError('Name');
    const AgeError = isFieldTouched('Age') && getFieldError('Age');
    const AddressError = isFieldTouched('Address') && getFieldError('Address');
    return (
      <Form layout="inline" onSubmit={this.handleSubmit}>
        <Form.Item validateStatus={NameError ? 'error' : ''} help={NameError || ''}>
          {getFieldDecorator('name', {
            rules: [{ required: true, message: 'Please input your Name!' }],
          })(
            <Input
              prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
              placeholder="Name"
            />,
          )}
        </Form.Item>
        <Form.Item validateStatus={AgeError ? 'error' : ''} help={AgeError || ''}>
          {getFieldDecorator('age', {
            rules: [{ required: true, message: 'Please input your Age!' }],
          })(
            <Input
              prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
              type="Age"
              placeholder="Age"
            />,
          )}
        </Form.Item>
        <Form.Item validateStatus={AddressError ? 'error' : ''} help={AddressError || ''}>
          {getFieldDecorator('address', {
            rules: [{ required: true, message: 'Please input your Address!' }],
          })(
            <Input
              prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
              placeholder="Address"
            />,
          )}
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" disabled={hasErrors(getFieldsError())}>
            Submit                               
          </Button>
        </Form.Item>
      </Form>
    );
  }
}

const WrappedHorizontalLoginForm = Form.create({ Name: 'Add_Form' })(AddForm);

export default WrappedHorizontalLoginForm;