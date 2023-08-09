CREATE OR REPLACE PROCEDURE p (p_lowercase  IN  VARCHAR2, 
                               po_uppercase OUT VARCHAR2)
IS
BEGIN
  po_uppercase := UPPER(p_lowercase);   
END p;
/