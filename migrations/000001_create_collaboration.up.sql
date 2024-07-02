CREATE TYPE collaboration_role AS ENUM ('owner', 'collaborator', 'viewer');
CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE collaborations (
                                collaborations_id uuid default gen_random_uuid() PRIMARY KEY,
                                composition_id uuid NOT NULL ,
                                user_id uuid NOT NULL ,
                                role collaboration_role DEFAULT 'collaborator',
                                joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE invitations (
                             invitations_id uuid default gen_random_uuid() PRIMARY KEY,
                             composition_id uuid NOT NULL ,
                             inviter_id uuid NOT NULL ,
                             invitee_id uuid NOT NULL ,
                             status invitation_status DEFAULT 'pending',
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             deleted_at BIGINT default 0
);

CREATE TABLE comments (
                          comments_id uuid default gen_random_uuid() PRIMARY KEY,
                          composition_id uuid NOT NULL ,
                          user_id uuid NOT NULL ,
                          content TEXT,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          deleted_at BIGINT default 0
);
