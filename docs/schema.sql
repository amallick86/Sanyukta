-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-01-10T11:29:30.133Z

CREATE TABLE "Gender" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "gender" varchar UNIQUE NOT NULL
);

CREATE TABLE "MaritalStatus" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "maritalstatus" varchar UNIQUE NOT NULL
);

CREATE TABLE "Religion" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "religion" varchar UNIQUE NOT NULL
);

CREATE TABLE "Caste" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "caste" varchar UNIQUE NOT NULL
);

CREATE TABLE "HigherEducation" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "education" varchar UNIQUE NOT NULL
);

CREATE TABLE "Profession" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "profession" varchar UNIQUE NOT NULL
);

CREATE TABLE "Country" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "country" varchar UNIQUE NOT NULL,
  "ISO2" varchar NOT NULL
);

CREATE TABLE "State" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "CountryID" integer,
  "state" varchar NOT NULL
);

CREATE TABLE "ResidingStatus" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "residingstatus" varchar UNIQUE NOT NULL
);

CREATE TABLE "Zodiac" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "zodiac" varchar UNIQUE NOT NULL
);

CREATE TABLE "MatchesStatus" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "status" varchar UNIQUE NOT NULL
);

CREATE TABLE "Features" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "feature" varchar UNIQUE NOT NULL
);

CREATE TABLE "SubscriptionTypes" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "Type" varchar UNIQUE NOT NULL
);

CREATE TABLE "Subscription" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "TypeID" integer,
  "Price" integer NOT NULL,
  "Duration" varchar NOT NULL
);

CREATE TABLE "SubscriptionFeatures" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "SubscriptionID" integer,
  "FeaturesID" integer
);

CREATE TABLE "User" (
  "UID" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "FirstName" varchar NOT NULL,
  "LastName" varchar NOT NULL,
  "Email" varchar UNIQUE NOT NULL,
  "IsVerified" bool NOT NULL DEFAULT false,
  "GoogelID" varchar UNIQUE NOT NULL,
  "SubscriptionID" integer DEFAULT 0,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "EmailOTP" (
  "ID" integer PRIMARY KEY,
  "UID" uuid,
  "OTP" integer NOT NULL,
  "CountOfDay" integer NOT NULL,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "ExpiredAt" timestamptz NOT NULL
);

CREATE TABLE "UserPhone" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "UID" uuid,
  "number" integer NOT NULL,
  "IsVerfied" bool NOT NULL DEFAULT false,
  "OTP" varchar NOT NULL,
  "NextOTP" timestamptz NOT NULL
);

CREATE TABLE "UserDetails" (
  "UDID" integer PRIMARY KEY,
  "UID" uuid,
  "Bio" text NOT NULL,
  "DateOfBirth" date NOT NULL,
  "Gender" integer,
  "Child" integer NOT NULL DEFAULT 0,
  "Siblings" integer NOT NULL DEFAULT 0,
  "Height" varchar NOT NULL,
  "Religion" integer,
  "Caste" integer,
  "Zodiac" integer,
  "MaritalStatus" integer,
  "HigherEducation" integer,
  "Profession" integer,
  "HomeCountry" integer,
  "HomeState" integer,
  "ResidingCountry" integer,
  "ResidingState" integer,
  "ResidingStatus" integer,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "UpdatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Sessions" (
  "ID" uuid PRIMARY KEY,
  "UID" uuid,
  "RefreshToken" varchar NOT NULL,
  "UserAgent" varchar NOT NULL,
  "ClientIP" varchar NOT NULL,
  "IsBlocked" bool NOT NULL DEFAULT false,
  "ExpiresAt" timestamptz NOT NULL,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "FCToken" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "UID" uuid,
  "UserAgent" varchar NOT NULL,
  "ClientIP" varchar NOT NULL,
  "Token" varchar NOT NULL,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Matches" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "SenderID" uuid,
  "ReceiverID" uuid,
  "StatusID" integer,
  "IsRead" bool NOT NULL DEFAULT false,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "UpdatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Block" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "UID" uuid,
  "BlockedID" uuid,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "UpdatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Favourite" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "UserID" uuid,
  "ProfileID" uuid,
  "Favourite" bool NOT NULL DEFAULT true,
  "IsRead" bool NOT NULL DEFAULT false,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "UpdatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "History" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "UserID" uuid,
  "ProfileID" uuid,
  "VisitedTime" integer NOT NULL,
  "IsRead" bool NOT NULL DEFAULT false,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "UpdatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "PaymentMethod" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "Method" varchar NOT NULL
);

CREATE TABLE "Payment" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "UID" uuid,
  "Method" integer,
  "SubscriptionID" integer,
  "AmountPaied" integer NOT NULL,
  "TxID" varchar NOT NULL,
  "ExpiresAt" timestamptz NOT NULL,
  "PaidAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Message" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "SenderID" uuid,
  "ReceiverID" uuid,
  "Data" text NOT NULL,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "UpdatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "MessageRead" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "MessageID" integer,
  "UID" uuid
);

CREATE TABLE "MessageDelete" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "MessageID" integer,
  "UID" uuid
);

CREATE TABLE "NotificationType" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "Type" varchar NOT NULL
);

CREATE TABLE "NotificationStatus" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "Status" varchar NOT NULL
);

CREATE TABLE "Notification" (
  "ID" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "UID" uuid,
  "ProfileID" uuid,
  "NotificationType" integer,
  "NotificationMessage" text,
  "IsRead" bool NOT NULL DEFAULT false,
  "NotificationStatus" integer,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "State" ADD FOREIGN KEY ("CountryID") REFERENCES "Country" ("ID");

ALTER TABLE "Subscription" ADD FOREIGN KEY ("TypeID") REFERENCES "SubscriptionTypes" ("ID");

ALTER TABLE "SubscriptionFeatures" ADD FOREIGN KEY ("SubscriptionID") REFERENCES "Subscription" ("ID");

ALTER TABLE "SubscriptionFeatures" ADD FOREIGN KEY ("FeaturesID") REFERENCES "Features" ("ID");

ALTER TABLE "User" ADD FOREIGN KEY ("SubscriptionID") REFERENCES "Subscription" ("ID");

ALTER TABLE "UserPhone" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("Gender") REFERENCES "Gender" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("Religion") REFERENCES "Religion" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("Caste") REFERENCES "Caste" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("Zodiac") REFERENCES "Zodiac" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("MaritalStatus") REFERENCES "MaritalStatus" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("HigherEducation") REFERENCES "HigherEducation" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("Profession") REFERENCES "Profession" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("HomeCountry") REFERENCES "Country" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("HomeState") REFERENCES "State" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("ResidingCountry") REFERENCES "Country" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("ResidingState") REFERENCES "State" ("ID");

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("ResidingStatus") REFERENCES "ResidingStatus" ("ID");

ALTER TABLE "Sessions" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "FCToken" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "Matches" ADD FOREIGN KEY ("SenderID") REFERENCES "User" ("UID");

ALTER TABLE "Matches" ADD FOREIGN KEY ("ReceiverID") REFERENCES "User" ("UID");

ALTER TABLE "Matches" ADD FOREIGN KEY ("StatusID") REFERENCES "MatchesStatus" ("ID");

ALTER TABLE "Block" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "Block" ADD FOREIGN KEY ("BlockedID") REFERENCES "User" ("UID");

ALTER TABLE "Favourite" ADD FOREIGN KEY ("UserID") REFERENCES "User" ("UID");

ALTER TABLE "Favourite" ADD FOREIGN KEY ("ProfileID") REFERENCES "User" ("UID");

ALTER TABLE "History" ADD FOREIGN KEY ("UserID") REFERENCES "User" ("UID");

ALTER TABLE "History" ADD FOREIGN KEY ("ProfileID") REFERENCES "User" ("UID");

ALTER TABLE "Payment" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "Payment" ADD FOREIGN KEY ("Method") REFERENCES "PaymentMethod" ("ID");

ALTER TABLE "Payment" ADD FOREIGN KEY ("SubscriptionID") REFERENCES "Subscription" ("ID");

ALTER TABLE "Message" ADD FOREIGN KEY ("SenderID") REFERENCES "User" ("UID");

ALTER TABLE "Message" ADD FOREIGN KEY ("ReceiverID") REFERENCES "User" ("UID");

ALTER TABLE "MessageRead" ADD FOREIGN KEY ("MessageID") REFERENCES "Message" ("ID");

ALTER TABLE "MessageRead" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "MessageDelete" ADD FOREIGN KEY ("MessageID") REFERENCES "Message" ("ID");

ALTER TABLE "MessageDelete" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "Notification" ADD FOREIGN KEY ("UID") REFERENCES "User" ("UID");

ALTER TABLE "Notification" ADD FOREIGN KEY ("ProfileID") REFERENCES "User" ("UID");

ALTER TABLE "Notification" ADD FOREIGN KEY ("NotificationType") REFERENCES "NotificationType" ("ID");

ALTER TABLE "Notification" ADD FOREIGN KEY ("NotificationStatus") REFERENCES "NotificationStatus" ("ID");
