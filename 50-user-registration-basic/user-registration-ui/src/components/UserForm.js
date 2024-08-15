// src/components/UserForm.js
import React, { useState } from 'react';

const UserForm = ({ onSubmit, user }) => {
    const [formData, setFormData] = useState(user || {
        first_name: '',
        last_name: '',
        gender: '',
        age: '',
        address: '',
        contact_no: '',
        contact_email: '',
        photo: ''
    });

    const [errors, setErrors] = useState({});

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const validate = () => {
        const newErrors = {};
        if (!formData.first_name) newErrors.first_name = 'First Name is required';
        if (!formData.last_name) newErrors.last_name = 'Last Name is required';
        if (!formData.gender) newErrors.gender = 'Gender is required';
        if (!formData.age) newErrors.age = 'Age is required';
        if (!formData.address) newErrors.address = 'Address is required';
        if (!formData.contact_no) newErrors.contact_no = 'Contact No is required';
        if (!formData.contact_email) newErrors.contact_email = 'Contact Email is required';
        if (!formData.photo) newErrors.photo = 'Photo URL is required';
        return newErrors;
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        const newErrors = validate();
        if (Object.keys(newErrors).length > 0) {
            setErrors(newErrors);
        } else {
            // onSubmit(formData);
            // Convert age to integer before submitting
            const dataToSubmit = { ...formData, age: parseInt(formData.age, 10) };
            onSubmit(dataToSubmit);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" name="first_name" value={formData.first_name} onChange={handleChange} placeholder="First Name" required />
            {errors.first_name && <span>{errors.first_name}</span>}
            <input type="text" name="last_name" value={formData.last_name} onChange={handleChange} placeholder="Last Name" required />
            {errors.last_name && <span>{errors.last_name}</span>}
            <input type="text" name="gender" value={formData.gender} onChange={handleChange} placeholder="Gender" required />
            {errors.gender && <span>{errors.gender}</span>}
            <input type="number" name="age" value={formData.age} onChange={handleChange} placeholder="Age" required />
            {errors.age && <span>{errors.age}</span>}
            <input type="text" name="address" value={formData.address} onChange={handleChange} placeholder="Address" required />
            {errors.address && <span>{errors.address}</span>}
            <input type="text" name="contact_no" value={formData.contact_no} onChange={handleChange} placeholder="Contact No" required />
            {errors.contact_no && <span>{errors.contact_no}</span>}
            <input type="email" name="contact_email" value={formData.contact_email} onChange={handleChange} placeholder="Contact Email" required />
            {errors.contact_email && <span>{errors.contact_email}</span>}
            <input type="text" name="photo" value={formData.photo} onChange={handleChange} placeholder="Photo URL" required />
            {errors.photo && <span>{errors.photo}</span>}
            <button type="submit">Submit</button>

        </form>
    );
};

export default UserForm;
