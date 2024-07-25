CREATE TABLE IF NOT EXISTS country
(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    country_name VARCHAR(255),
    created TIMESTAMP
);

CREATE TABLE IF NOT EXISTS engineer
(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    gender SMALLINT NOT NULL,
    country_id BIGINT,
    title VARCHAR(255),
    created TIMESTAMP,
    FOREIGN KEY (country_id) REFERENCES country(id)
);

insert into country(country_name, created) values ('Afghanistan', current_timestamp);
insert into country(country_name, created) values ('Albania', current_timestamp);
insert into country(country_name, created) values ('Algeria', current_timestamp);
insert into country(country_name, created) values ('American Samoa', current_timestamp);
insert into country(country_name, created) values ('Andorra', current_timestamp);
insert into country(country_name, created) values ('Angola', current_timestamp);
insert into country(country_name, created) values ('Anguilla', current_timestamp);
insert into country(country_name, created) values ('Antarctica', current_timestamp);
insert into country(country_name, created) values ('Antigua and Barbuda', current_timestamp);
insert into country(country_name, created) values ('Argentina', current_timestamp);
insert into country(country_name, created) values ('Armenia', current_timestamp);
insert into country(country_name, created) values ('Aruba', current_timestamp);
insert into country(country_name, created) values ('Australia', current_timestamp);
insert into country(country_name, created) values ('Austria', current_timestamp);
insert into country(country_name, created) values ('Azerbaijan', current_timestamp);
insert into country(country_name, created) values ('Bahamas', current_timestamp);
insert into country(country_name, created) values ('Bahrain', current_timestamp);
insert into country(country_name, created) values ('Bangladesh', current_timestamp);
insert into country(country_name, created) values ('Barbados', current_timestamp);
insert into country(country_name, created) values ('Belarus', current_timestamp);
insert into country(country_name, created) values ('Belgium', current_timestamp);
insert into country(country_name, created) values ('Belize', current_timestamp);
insert into country(country_name, created) values ('Benin', current_timestamp);
insert into country(country_name, created) values ('Bermuda', current_timestamp);
insert into country(country_name, created) values ('Bhutan', current_timestamp);
insert into country(country_name, created) values ('Bolivia', current_timestamp);
insert into country(country_name, created) values ('Bosnia and Herzegowina', current_timestamp);
insert into country(country_name, created) values ('Botswana', current_timestamp);
insert into country(country_name, created) values ('Bouvet Island', current_timestamp);
insert into country(country_name, created) values ('Brazil', current_timestamp);
insert into country(country_name, created) values ('British Indian Ocean Territory', current_timestamp);
insert into country(country_name, created) values ('Brunei Darussalam', current_timestamp);
insert into country(country_name, created) values ('Bulgaria', current_timestamp);
insert into country(country_name, created) values ('Burkina Faso', current_timestamp);
insert into country(country_name, created) values ('Burundi', current_timestamp);
insert into country(country_name, created) values ('Cambodia', current_timestamp);
insert into country(country_name, created) values ('Cameroon', current_timestamp);
insert into country(country_name, created) values ('Canada', current_timestamp);
insert into country(country_name, created) values ('Cape Verde', current_timestamp);
insert into country(country_name, created) values ('Cayman Islands', current_timestamp);
insert into country(country_name, created) values ('Central African Republic', current_timestamp);
insert into country(country_name, created) values ('Chad', current_timestamp);
insert into country(country_name, created) values ('Chile', current_timestamp);
insert into country(country_name, created) values ('China', current_timestamp);
insert into country(country_name, created) values ('Christmas Island', current_timestamp);
insert into country(country_name, created) values ('Cocos (Keeling) Islands', current_timestamp);
insert into country(country_name, created) values ('Colombia', current_timestamp);
insert into country(country_name, created) values ('Comoros', current_timestamp);
insert into country(country_name, created) values ('Congo', current_timestamp);
insert into country(country_name, created) values ('Congo, the Democratic Republic of the', current_timestamp);
insert into country(country_name, created) values ('Cook Islands', current_timestamp);
insert into country(country_name, created) values ('Costa Rica', current_timestamp);
insert into country(country_name, created) values ('Cote d Ivoire', current_timestamp);
insert into country(country_name, created) values ('Croatia (Hrvatska)', current_timestamp);
insert into country(country_name, created) values ('Cuba', current_timestamp);
insert into country(country_name, created) values ('Cyprus', current_timestamp);
insert into country(country_name, created) values ('Czech Republic', current_timestamp);
insert into country(country_name, created) values ('Denmark', current_timestamp);
insert into country(country_name, created) values ('Djibouti', current_timestamp);
insert into country(country_name, created) values ('Dominica', current_timestamp);
insert into country(country_name, created) values ('Dominican Republic', current_timestamp);
insert into country(country_name, created) values ('East Timor', current_timestamp);
insert into country(country_name, created) values ('Ecuador', current_timestamp);
insert into country(country_name, created) values ('Egypt', current_timestamp);
insert into country(country_name, created) values ('El Salvador', current_timestamp);
insert into country(country_name, created) values ('Equatorial Guinea', current_timestamp);
insert into country(country_name, created) values ('Eritrea', current_timestamp);
insert into country(country_name, created) values ('Estonia', current_timestamp);
insert into country(country_name, created) values ('Ethiopia', current_timestamp);
insert into country(country_name, created) values ('Falkland Islands (Malvinas)', current_timestamp);
insert into country(country_name, created) values ('Faroe Islands', current_timestamp);
insert into country(country_name, created) values ('Fiji', current_timestamp);
insert into country(country_name, created) values ('Finland', current_timestamp);
insert into country(country_name, created) values ('France', current_timestamp);
insert into country(country_name, created) values ('France Metropolitan', current_timestamp);
insert into country(country_name, created) values ('French Guiana', current_timestamp);
insert into country(country_name, created) values ('French Polynesia', current_timestamp);
insert into country(country_name, created) values ('French Southern Territories', current_timestamp);
insert into country(country_name, created) values ('Gabon', current_timestamp);
insert into country(country_name, created) values ('Gambia', current_timestamp);
insert into country(country_name, created) values ('Georgia', current_timestamp);
insert into country(country_name, created) values ('Germany', current_timestamp);
insert into country(country_name, created) values ('Ghana', current_timestamp);
insert into country(country_name, created) values ('Gibraltar', current_timestamp);
insert into country(country_name, created) values ('Greece', current_timestamp);
insert into country(country_name, created) values ('Greenland', current_timestamp);
insert into country(country_name, created) values ('Grenada', current_timestamp);
insert into country(country_name, created) values ('Guadeloupe', current_timestamp);
insert into country(country_name, created) values ('Guam', current_timestamp);
insert into country(country_name, created) values ('Guatemala', current_timestamp);
insert into country(country_name, created) values ('Guinea', current_timestamp);
insert into country(country_name, created) values ('Guinea-Bissau', current_timestamp);
insert into country(country_name, created) values ('Guyana', current_timestamp);
insert into country(country_name, created) values ('Haiti', current_timestamp);
insert into country(country_name, created) values ('Heard and Mc Donald Islands', current_timestamp);
insert into country(country_name, created) values ('Holy See (Vatican City State)', current_timestamp);
insert into country(country_name, created) values ('Honduras', current_timestamp);
insert into country(country_name, created) values ('Hong Kong', current_timestamp);
insert into country(country_name, created) values ('Hungary', current_timestamp);
insert into country(country_name, created) values ('Iceland', current_timestamp);
insert into country(country_name, created) values ('India', current_timestamp);